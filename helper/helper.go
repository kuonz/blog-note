package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
)

var exp1 = regexp.MustCompile(`(?U)!\[.*\]\(.*\)`)
var exp2 = regexp.MustCompile(`(?U)\[.*\]`)
var exp3 = regexp.MustCompile(`(?U)\(.*\)`)

var staticPath = ""
var processCount = 0

func main() {
	fmt.Println("将处理缩进、头部信息、图片地址")

	traversal("..", "")

	fmt.Printf("process %d article.\n", processCount)
}

func traversal(prefix string, fileName string) {

	currentPath := prefix + fileName + "/"

	infos, _ := ioutil.ReadDir(currentPath)

	for _, item := range infos {
		itemName := item.Name()

		if item.IsDir() {
			traversal(currentPath, itemName)
		} else {
			if strings.HasSuffix(itemName, ".md") {
				titleName := getTitleName(itemName)
				processMD(currentPath, itemName, titleName)
			}
		}
	}
}

func processMD(filePath string, fileName string, titleName string) {

	contentByte, _ := ioutil.ReadFile(filePath + "/" + fileName)
	contentString := string(contentByte)

	if !strings.HasPrefix(contentString, `---`) {
		staticPath = strings.ReplaceAll(filePath, "..", "(/post")

		currentDate := time.Now().Format("2006-01-02")

		temp := strings.Split(filePath, "/")
		currentCatogories := getTitleName(temp[len(temp)-2])

		processContent(&contentString, titleName, currentDate, currentCatogories)

		ioutil.WriteFile(filePath+"/"+fileName, []byte(contentString), 0666)

		fmt.Printf("processed %s ... \n", filePath + "/" + fileName)
		processCount++
	}
}

func processContent(src *string, title, date, categories string) {

	// 添加头部信息
	tl := `---
author: "kuonz"
draft: false
title: "` + title + `"
date: ` + date + `
categories: ["` + categories + `"]
---
	
`
	*src = tl + *src

	// 处理缩进
	*src = strings.ReplaceAll(*src, "\t", "  ")

	// 处理图片
	*src = exp1.ReplaceAllStringFunc(*src, processImage)
}

func processImage(src string) string {
	return exp3.ReplaceAllStringFunc(exp2.ReplaceAllString(src, "[]"), processImage2)
}

func processImage2(src string) string {
	return strings.ReplaceAll(src, "(", staticPath)
}

func getTitleName(fileName string) string {
	index := strings.Index(fileName, "-")
	if index > 0 {
		fileName = fileName[index+1:]
	}
	index = strings.LastIndex(fileName, ".md")
	if index > 0 {
		fileName = fileName[:index]
	}
	return fileName
}
