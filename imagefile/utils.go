package imageFile
import (
  "os"
  "image"
  "path/filepath"
  "log"
)

func absPath(path string) (apath string) {
  apath, err := filepath.Abs(path)
  logError(err)
  return
}

func stat(file *os.File) (fi os.FileInfo) {
  fi, err := file.Stat()
  logError(err)
  return
}

func openFile(path string) (f *os.File) {
  f, err := os.Open(path)
  logError(err)
  return
}

func decodeConfig(file *os.File) (cfg image.Config, format string) {
  cfg, format, err := image.DecodeConfig(file)
  logError(err)
  return
}

func logError(e error) {
  if e != nil {
    log.Fatal(e)
  }
}
