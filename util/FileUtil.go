package util

func isDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}

	panic("not reached")
}

/** 复制文件到指定地方 **/
func CopyFile(src, dest string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	defer srcFile.Close()

	desFile, err := os.Create(dest)
	if err != nil {
		fmt.Println(err)
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}

/**复制某个文件夹的文件到另一个文件夹 **/
func CopyDir(srcDir, destDir string) {

	if isDirExists(srcDir) {
		//存在附件文件夹
		tmpSrc := strings.TrimSpace(srcDir)
		files, _ := ioutil.ReadDir(tmpSrc)
		for _, f := range files {
			CopyFile(srcDir+f.Name(), destDir+f.Name())
		}
	}
}

func joinStr(lines []string) {
	result := bytes.Buffer{}
	for i := end; i < len(lines); i++ {
		result.WriteString(lines[i])
	}

	return result.String()
}
