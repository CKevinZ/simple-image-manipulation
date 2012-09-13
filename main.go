package main
import (
  . "./imagefile"
  . "fmt"
  "log"
  "os"
  "flag"
  "strconv"
  "strings"
  "io/ioutil"
  "encoding/json"
  "image"
)

func info(i *ImageFile) {
  jsonbytes,_ := json.Marshal(i)
  Print(string(jsonbytes))
}

func crop(i *ImageFile, geo geometry, out, format string) {
  if len(geo) != 4 {
    log.Fatal("-geo arguments malformed (usage: 'x0 y0 x1 y1')")
  }
  cropedImg := i.Crop(geo[0], geo[1], geo[2], geo[3])
  Print(writeImage(cropedImg, out, format))
}

func resize(i *ImageFile, geo geometry, out, format string) {
  if len(geo) != 2 {
    log.Fatal("-geo arguments malformed (usage: 'x y')")
  }
  resizedImage := i.Resize(geo[0], geo[1])
  Print(writeImage(resizedImage, out, format))
}

type geometry []int

func (g *geometry) String() string {
  return Sprint(*g)
}

func (g *geometry) Set(value string) error {
  for _, str := range strings.Split(value, " ") {
    i, err := strconv.ParseInt(str, 10, 0)
    if err != nil {
      return err
    }
    *g = append(*g, int(i))
  }
  return nil
}

func main() {
  var (
    i, c, r, o, f string
    g geometry
  )

  flag.StringVar(&i, "info", "", "Return a JSON containing an image file informations.")
  flag.StringVar(&c, "crop", "", "Crop an image file.")
  flag.StringVar(&r, "resize", "", "Resize an image file.")
  flag.Var(&g, "geo", "Geometry parameters (usage: 'x0 y0 x1 y1').")
  flag.StringVar(&o, "out", "", "Output file path (none: /tmp/image_#rand.")
  flag.StringVar(&f, "format", "jpg", "Output file format (jpg or png default is jpg).")
  flag.Parse()

  switch {
  case i != "": info(NewImageFile(i))
  case c != "": crop(NewImageFile(c), g, o, f)
  case r != "": resize(NewImageFile(r), g, o, f)
  default: log.Fatal("Argument Error: type -h for help")
  }
}

func newFile(path string) (file *os.File) {
  file, err := os.Create(path)
  if err != nil {
    log.Fatal(err)
  }
  return
}

func newTmpFile() (tmpFile *os.File) {
  tmpFile, err := ioutil.TempFile("", "image_")
  if err != nil {
    log.Fatal(err)
  }
  return
}

func writeImage(img image.Image, out, format string) string {
  var file *os.File
  if out == "" {
    file = newTmpFile()
  } else {
    file = newFile(out)
  }
  Encode(file, img, format)
  file.Close()
  return file.Name()
}
