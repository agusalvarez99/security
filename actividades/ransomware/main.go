package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	menu()
}

func encrypt() {
	//va a leer los archivos de la carpeta cifrar
	//va a guardar los archivos enciptados en la carpeta cifrados
	listaFicheros, err := ioutil.ReadDir("./cifrar")
	if err != nil {
		log.Fatal(err)
	}
	publicKey, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	if err != nil {
		log.Fatal(err)
	}
	bloque, err := aes.NewCipher(publicKey)
	if err != nil {
		log.Panic(err)
	}
	for _, f := range listaFicheros {
		contenido, err := ioutil.ReadFile("./cifrar/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		gcm, err := cipher.NewGCM(bloque)
		if err != nil {
			log.Panic(err)
		}
		nonce := make([]byte, gcm.NonceSize())
		cifrado := gcm.Seal(nonce, nonce, contenido, nil)
		fmt.Println("Archivo encriptado:", f.Name())
		err = ioutil.WriteFile("./cifrados/"+f.Name()+"Cifrado.bin", cifrado, 0777)
		if err != nil {
			log.Panic(err)
		}

	}
}

func decrypt() {
	//va a leer del directorio cifrados
	//va a dejar los archivos descifrados en la carpeta actual
	listaFicheros, err := ioutil.ReadDir("./cifrados")
	if err != nil {
		log.Fatal(err)
	}
	publicKey, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	if err != nil {
		log.Fatal(err)
	}
	bloque, err := aes.NewCipher(publicKey)
	if err != nil {
		log.Panic(err)
	}
	for _, f := range listaFicheros {
		contenido, err := ioutil.ReadFile("./cifrados/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		gcm, err := cipher.NewGCM(bloque)
		if err != nil {
			log.Panic(err)
		}
		nonce := contenido[:gcm.NonceSize()]
		contenido = contenido[gcm.NonceSize():]
		cifrado, err := gcm.Open(nil, nonce, contenido, nil)
		if err != nil {
			log.Panic(err)
		}
		err = ioutil.WriteFile(f.Name()+"Descifrado.txt", cifrado, 0777)
		if err != nil {
			log.Panic(err)
		}

	}
}

func menu() {
	opcion := 0
	for opcion != 3 {
		fmt.Println(" ")
		fmt.Println("Opciones")
		fmt.Println("1 - Encriptar archivos")
		fmt.Println("2 - Desencriptar archivos")
		fmt.Println("3 - Salir")
		fmt.Printf("Ingrese la opcion que desee: ")
		fmt.Scanln(&opcion)
		if opcion == 1 {
			encrypt()
		} else if opcion == 2 {
			decrypt()
		} else if opcion != 3 {
			fmt.Println("Opcion no valida!")
		}
	}
}
