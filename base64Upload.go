package base64Upload

import (
      "os"
      "io"
      "encoding/base64"
)

func Upload(fileName string, content string) {
      decode, err := base64.StdEncoding.DecodeString(content)
      if err != nil {
            return err
      }
      file, err := os.Create(fileName)
      if err != nil {
            return err
      }
      defer file.Close()
      _, err := file.Write(decode)
      if err != nil {
            return err
      }
}
