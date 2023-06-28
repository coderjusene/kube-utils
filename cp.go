package kube_utils

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	v1core "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"kube-center/kube-utils/client"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func (k *Pod) CopyToPod(src, dest string) error {
	restconfig, err, coreclient := client.InitRestClient()
	if err != nil {
		return err
	}

	reader, writer := io.Pipe()

	go func() {
		defer writer.Close()
		cmdutil.CheckErr(makeTar(src, dest, writer))
	}()

	req := coreclient.RESTClient().Post().
		Resource("pods").
		Name(k.Name).
		Namespace(k.Namespace).
		SubResource("exec").
		VersionedParams(&v1core.PodExecOptions{
			Stdin:   true,
			Stdout:  true,
			Stderr:  true,
			TTY:     false,
			Command: []string{"tar", "-xmf", "-"},
		}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(restconfig, "POST", req.URL())
	if err != nil {
		return err
	}

	err = exec.StreamWithContext(context.TODO(), remotecommand.StreamOptions{
		Stdin:  reader,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Tty:    false,
	})

	if err != nil {
		return err
	}

	return err
}

func (k *Pod) CopyFromPod(src, dest string) error {
	restconfig, err, coreclient := client.InitRestClient()
	if err != nil {
		return err
	}

	reader, writer := io.Pipe()

	req := coreclient.RESTClient().Get().
		Resource("pods").
		Name(k.Name).
		Namespace(k.Namespace).
		SubResource("exec").
		VersionedParams(&v1core.PodExecOptions{
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       false,
			Container: k.ContainerName,
			Command:   []string{"tar", "cf", "-", src},
		}, scheme.ParameterCodec)

	// remotecommand 主要实现了http 转 SPDY 添加X-Stream-Protocol-Version相关header 并发送请求
	exec, err := remotecommand.NewSPDYExecutor(restconfig, "POST", req.URL())
	if err != nil {
		return err
	}

	go func() {
		defer writer.Close()
		err := exec.StreamWithContext(context.TODO(), remotecommand.StreamOptions{
			Stdin:  os.Stdin,
			Stdout: writer,
			Stderr: os.Stderr,
			Tty:    false,
		})
		cmdutil.CheckErr(err)
	}()

	prefix := getPrefix(dest)
	prefix = path.Clean(prefix)
	prefix = stripPathShortcuts(prefix)
	destPath := path.Join("./", path.Base(prefix))
	err = unTarAll(reader, destPath, prefix)

	return err
}

func makeTar(src, dest string, writer io.Writer) error {
	// TODO: use compression here?
	gzipWriter := gzip.NewWriter(writer)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	// 清理 . 和 .. path
	src = path.Clean(src)
	dest = path.Clean(dest)

	// 递归tar
	return recursiverTar(path.Dir(src), path.Base(src), path.Dir(dest), path.Base(dest), tarWriter)
}

func recursiverTar(srcBase, srcFile, destBase, destFile string, tw *tar.Writer) error {
	srcPath := path.Join(srcBase, srcFile)
	matchedPath, err := filepath.Glob(srcPath)
	if err != nil {
		return err
	}

	for _, fpath := range matchedPath {
		stat, err := os.Lstat(fpath)
		if err != nil {
			return err
		}

		if stat.IsDir() {
			files, err := ioutil.ReadDir(fpath)
			if err != nil {
				return err
			}

			if len(files) == 0 {
				// case empty directory
				hdr, _ := tar.FileInfoHeader(stat, fpath)
				hdr.Name = destFile
				if err := tw.WriteHeader(hdr); err != nil {
					return err
				}
			}

			for _, f := range files {
				if err := recursiverTar(srcBase, path.Join(srcFile, f.Name()), destBase, path.Join(destFile, f.Name()), tw); err != nil {
					return err
				}
			}

			return nil
		} else if stat.Mode()&os.ModeSymlink != 0 {
			// case soft link
			hdr, _ := tar.FileInfoHeader(stat, fpath)
			target, err := os.Readlink(fpath)
			if err != nil {
				return err
			}
			hdr.Linkname = target
			hdr.Name = destFile
			if err := tw.WriteHeader(hdr); err != nil {
				return err
			}
		} else {
			// case regular file or other file type link pipe
			hdr, err := tar.FileInfoHeader(stat, fpath)
			if err != nil {
				return err
			}
			hdr.Name = destFile

			if err := tw.WriteHeader(hdr); err != nil {
				return err
			}

			f, err := os.Open(fpath)
			if err != nil {
				return err
			}
			defer f.Close()

			if _, err := io.Copy(tw, f); err != nil {
				return err
			}
			return f.Close()
		}
	}
	return nil
}

func getPrefix(file string) string {
	return strings.TrimLeft(file, "/")
}

// stripPathShortcuts removes any leading or trailing "../" from a given path
func stripPathShortcuts(p string) string {

	newPath := path.Clean(p)
	trimmed := strings.TrimPrefix(newPath, "../")

	for trimmed != newPath {
		newPath = trimmed
		trimmed = strings.TrimPrefix(newPath, "../")
	}

	// trim leftover {".", ".."}
	if newPath == "." || newPath == ".." {
		newPath = ""
	}

	if len(newPath) > 0 && string(newPath[0]) == "/" {
		return newPath[1:]
	}

	return newPath
}

func unTarAll(reader io.Reader, destDir, prefix string) error {
	tarReader := tar.NewReader(reader)
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		if !strings.HasPrefix(header.Name, prefix) {
			return fmt.Errorf("tar contents corrupted")
		}

		mode := header.FileInfo().Mode()
		destFileName := filepath.Join(destDir, header.Name[len(prefix):])

		baseName := filepath.Dir(destFileName)
		if err := os.MkdirAll(baseName, 0755); err != nil {
			return err
		}
		if header.FileInfo().IsDir() {
			if err := os.MkdirAll(destFileName, 0755); err != nil {
				return err
			}
			continue
		}

		evaledPath, err := filepath.EvalSymlinks(baseName)
		if err != nil {
			return err
		}

		if mode&os.ModeSymlink != 0 {
			linkname := header.Linkname

			if !filepath.IsAbs(linkname) {
				_ = filepath.Join(evaledPath, linkname)
			}

			if err := os.Symlink(linkname, destFileName); err != nil {
				return err
			}
		} else {
			outFile, err := os.Create(destFileName)
			if err != nil {
				return err
			}
			defer outFile.Close()
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return err
			}
			if err := outFile.Close(); err != nil {
				return err
			}
		}
	}

	return nil
}
