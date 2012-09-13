package imageFile
import (
  "image"
  "fmt"
  "os"
  "io"
  "log"
  "image/jpeg"
  "image/png"
  _ "image/gif"
)

type ImageFile struct {
  Name, AbsPath, Format string
  Width, Height int
  Size   int64
  Stat   os.FileInfo
}

func Encode(iw io.Writer, img image.Image, format string) {
  switch format {
  case "jpg": jpeg.Encode(iw, img, nil)
  case "png": png.Encode(iw, img)
  default: log.Fatal("Argument Error: available formats are jpg and png")
  }
}

func (i *ImageFile) Crop(x0, y0, x1, y1 int) (rgbaImg *image.RGBA) {
  img := i.decode()
  rgbaImg = image.NewRGBA(image.Rect(0, 0, x1-x0, y1-y0))
  bounds := rgbaImg.Bounds()

  for x := 0; x < bounds.Dx(); x++ {
    for y := 0; y < bounds.Dy(); y++ {
      rgbaImg.Set(x, y, img.At(x0+x, y0+y))
    }
  }
  return
}

func (i *ImageFile) Resize(width, height int) (rgbaImg *image.RGBA) {
  img := i.decode()
  bounds := img.Bounds()
  rgbaImg = image.NewRGBA(image.Rect(0, 0, width, height))

  pacex := float64(bounds.Dx())/float64(width)
  pacey := float64(bounds.Dy())/float64(height)

  for x := 0; x < width; x++ {
    for y := 0; y < height; y++ {
      rgbaImg.Set(x, y, img.At(int(float64(x)*pacex), int(float64(y)*pacey)))
    }
  }
  return
}

func (i *ImageFile) MarshalJSON() ([]byte, error) {
  s := fmt.Sprintf(`{
    "bytes": %d,
    "filename": %q,
    "format": %q,
    "height": %d,
    "path": %q,
    "pixels": %d,
    "size": [%d,%d],
    "width": %d
  }`, i.Size, i.Name, i.Format, i.Height, i.AbsPath,
      i.Height*i.Width, i.Height, i.Width, i.Width)
  return []byte(s), nil
}

func (i *ImageFile) decode() image.Image {
  file        := openFile(i.AbsPath)
  img, _, err := image.Decode(file)

  logError(err)
  file.Close()

  return img
}

func NewImageFile(path string) *ImageFile {
  apath       := absPath(path)
  file        := openFile(apath)
  defer file.Close()
  cfg, format := decodeConfig(file)
  stat        := stat(file)

  return &ImageFile{
    AbsPath: apath,
    Format:  format,
    Stat:    stat,
    Name:    stat.Name(),
    Size:    stat.Size(),
    Width:   cfg.Width,
    Height:  cfg.Height,
  }
}
