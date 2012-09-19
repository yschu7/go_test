// Calculate the CRC32/SHA1 checksum of the input file
package main

import (
  "fmt"
  "flag"
  "os"
  "hash/crc32"
  "crypto/sha1"
  "io/ioutil"
)

func usage() {
  fmt.Fprintf(os.Stderr, "usage: %s -crc|-sha inputfile\n", os.Args[0])
  flag.PrintDefaults()
  os.Exit(2)
}

func getCRC32(filename string) (uint32, error) {
  bs, err := ioutil.ReadFile(filename)
  if err != nil {
    return 0, err
  }
  h := crc32.NewIEEE()
  h.Write(bs)
  return h.Sum32(), nil
}

func getSHA1(filename string) ([]byte, error) {
  bs, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  h := sha1.New()
  h.Write([]byte(bs))
  return h.Sum([]byte{}), nil
}

func main() {
  flag.Usage = usage
  crc := flag.Bool("crc", false, "CRC32 flag")
  sha := flag.Bool("sha", false, "SHA1 flag")
  flag.Parse()
  args := flag.Args()
  if len(args) < 1 {
    printErr("Input file is missing.")
  }
  infname := args[0]
  if *crc {
    fmt.Println("*crc")
    h1, err := getCRC32(infname)
    if err != nil {
      printErr(fmt.Sprintf("Input file: %s is not found!!",infname))
    }
    fmt.Printf("%x\n",h1)
  } else if *sha {
    fmt.Println("*sha")
    h1, err := getSHA1(infname)
    if err != nil {
      printErr(fmt.Sprintf("Input file: %s is not found!!",infname))
    }
    printHex(h1)
    fmt.Println("\t", args[0])
  } else {
    usage()
  }
}

func printHex(arr []byte) {
  for _, v := range arr {
    fmt.Printf("%02x",int(v))
  }
}

func printErr(msg string) {
  fmt.Println(msg)
  os.Exit(1)
}
