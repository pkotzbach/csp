package algorithm

import (
	"bytes"
	"os"
)

const defaultVal = -1

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func clearDomains(fields []*Field) {
	for _, field := range fields {
		for i := 0; i < len(field.domainAv); i++ {
			field.domainAv[i] = -1
		}
	}
}

func importBinary(path string) ([]*Field, int) {
	dat, err := os.ReadFile(path)
	check(err)
	rows := bytes.Split(dat, []byte("\n"))
	size := len(rows[0]) - 1 // -1, because there is \n
	result := make([]*Field, size*size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			field := rows[i][j]
			newField := Field{domainAv: make([]int, 2)}
			if field == '0' || field == '1' {
				newField.val = int(field) - '0'
				newField.constant = true
			} else if field != '\n' {
				newField.val = defaultVal
				newField.constant = false
			}
			result[i*size+j] = &newField
		}
	}

	clearDomains(result)
	return result, size
}

func importFuto(path string) ([]*Field, []inequality, int) {
	dat, err := os.ReadFile(path)
	check(err)
	rows := bytes.Split(dat, []byte("\n"))

	//deleting \n
	for i := 0; i < len(rows)-1; i++ {
		rows[i] = rows[i][:len(rows[i])-1]
	}

	fullSize := len(rows[0])
	size := fullSize/2 + 1
	result := make([]*Field, size*size)
	var inequalities []inequality

	for i := 0; i < fullSize; i++ {
		posI := i / 2
		temp := fullSize
		if i%2 == 1 {
			temp = size
		}
		for j := 0; j < temp; j++ {
			posJ := j / 2
			field := rows[i][j]
			newField := Field{domainAv: make([]int, size+1)}
			if field >= '1' && field <= '9' {
				newField.val = int(field) - '0'
				newField.constant = true
				result[size*posI+posJ] = &newField
			} else if field == 'x' {
				newField.val = defaultVal
				newField.constant = false
				result[size*posI+posJ] = &newField
			} else if i%2 == 0 {
				if field == '>' {
					inequalities = append(inequalities, inequality{lower: posI*size + posJ + 1, bigger: posI*size + posJ})
				} else if field == '<' {
					inequalities = append(inequalities, inequality{bigger: posI*size + posJ + 1, lower: posI*size + posJ})
				}
			} else if i%2 == 1 {
				if field == '<' {
					inequalities = append(inequalities, inequality{lower: posI*size + j, bigger: (posI+1)*size + j})
				} else if field == '>' {
					inequalities = append(inequalities, inequality{bigger: posI*size + j, lower: (posI+1)*size + j})
				}
			}
		}
	}

	clearDomains(result)
	return result, inequalities, size
}
