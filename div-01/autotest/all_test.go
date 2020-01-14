package tests

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	student ".."
	"github.com/mgutz/ansi"
)

var errorsArray []string

// BTreeIsEqual function
func BTreeIsEqual(rootA, rootB *student.TreeNode) bool {
	if (rootA == nil) && (rootB == nil) {
		return true
	}
	if ((rootA == nil) && (rootB != nil)) || ((rootA != nil) && (rootB == nil)) {
		return false
	}
	if rootA.Data != rootB.Data {
		return false
	}
	if (rootA.Data == rootB.Data) && ((rootA.Parent != nil) && (rootB.Parent != nil)) {
		if rootA.Parent.Data != rootB.Parent.Data {
			return false
		}
	}
	return (BTreeIsEqual(rootA.Left, rootB.Left) && BTreeIsEqual(rootA.Right, rootB.Right))
}

// BTreeSearchItem function
func BTreeSearchItem(root *student.TreeNode, dataref string) *student.TreeNode {
	if root == nil {
		return root
	}
	if dataref > root.Data {
		return BTreeSearchItem(root.Right, dataref)
	} else if dataref < root.Data {
		return BTreeSearchItem(root.Left, dataref)
	} else {
		return root
	}
}

// BTreeInsertData function
func BTreeInsertData(root *student.TreeNode, dataref string) *student.TreeNode {
	if root == nil {
		return &student.TreeNode{Data: dataref}
	}
	if dataref >= root.Data {
		root.Right = BTreeInsertData(root.Right, dataref)
		root.Right.Parent = root
	} else {
		root.Left = BTreeInsertData(root.Left, dataref)
		root.Left.Parent = root
	}
	return root
}

func printBTree(root *student.TreeNode) string {
	var result string
	printTreeRight(root, &result, "", "")
	return result
}

func printTreeRight(root *student.TreeNode, buffer *string, prefix string, childrenPrefix string) {
	if root != nil {
		*buffer += prefix
		pData := ""
		if root.Parent != nil {
			pData = root.Parent.Data
		} else {
			pData = "nil"
		}
		*buffer += fmt.Sprintf("%02s (%02s)", root.Data, pData)
		*buffer += "\n"
		if root.Right != nil {
			if root.Left != nil {
				printTreeRight(root.Right, buffer, childrenPrefix+"├── ", childrenPrefix+"│   ")
			} else {
				printTreeRight(root.Right, buffer, childrenPrefix+"└── ", childrenPrefix+"   ")
			}
		}
		if root.Left != nil {
			if root.Right != nil {
				printTreeRight(root.Left, buffer, childrenPrefix+"└── ", childrenPrefix+"    ")
			} else {
				printTreeRight(root.Left, buffer, childrenPrefix+"└── ", childrenPrefix+"    ")
			}
		}
	}
}

// PrintList2 func
func PrintList2(l *student.List) string {
	var result string
	it := l.Head
	for it != nil {
		result += fmt.Sprint(it.Data, " -> ")
		it = it.Next
	}
	result += fmt.Sprint(nil, "")
	return result
}

// PrintList function
func PrintList(l *student.NodeI) string {
	var result string
	it := l
	for it != nil {
		result += fmt.Sprint(it.Data, " -> ")
		it = it.Next
	}
	result += fmt.Sprint(nil, "")
	return result
}

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

// ListPushBack function
func ListPushBack(l *student.List, dataref interface{}) *student.List {
	var newNode = &student.NodeL{Data: dataref}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return l
	}
	l.Tail.Next = newNode
	l.Tail = l.Tail.Next
	return l
}

func status(pass bool) string {
	var result string
	if !pass {
		result = ansi.ColorCode("red+bold") + "FAIL" + ansi.ColorCode("reset")
	} else {
		result = ansi.ColorCode("green+bold") + "PASS" + ansi.ColorCode("reset")
	}
	return result
}

func dieOn(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

func catchPrintCombN(t *testing.T, runnable func(int), nbr int) string {
	realStdout := os.Stdout
	defer func() { os.Stdout = realStdout }()
	r, fakeStdout, err := os.Pipe()
	dieOn(err, t)
	os.Stdout = fakeStdout
	runnable(nbr)
	dieOn(fakeStdout.Close(), t)
	newOutBytes, err := ioutil.ReadAll(r)
	dieOn(err, t)
	dieOn(r.Close(), t)
	return strings.TrimSuffix(string(newOutBytes), "\n")
}

func catchPrintNbrBase(t *testing.T, runnable func(int, string), nbr int, base string) string {
	realStdout := os.Stdout
	defer func() { os.Stdout = realStdout }()
	r, fakeStdout, err := os.Pipe()
	dieOn(err, t)
	os.Stdout = fakeStdout
	runnable(nbr, base)
	dieOn(fakeStdout.Close(), t)
	newOutBytes, err := ioutil.ReadAll(r)
	dieOn(err, t)
	dieOn(r.Close(), t)
	return strings.TrimSuffix(string(newOutBytes), "\n")
}

func catchDoop(t *testing.T, program, a, op, b string) string {
	r, fakeStdout, err := os.Pipe()
	dieOn(err, t)
	cmd := exec.Command(program, a, op, b)
	cmd.Stdout = fakeStdout
	errRun := cmd.Run()
	if errRun != nil {
		dieOn(errRun, t)
	}
	dieOn(fakeStdout.Close(), t)
	newOutBytes, err := ioutil.ReadAll(r)
	dieOn(err, t)
	dieOn(r.Close(), t)
	return strings.TrimSuffix(string(newOutBytes), "\n")
}

func catchRotateVowels(t *testing.T, program string, input []string) string {
	r, fakeStdout, err := os.Pipe()
	dieOn(err, t)
	cmd := exec.Command(program, strings.Join(input, " "))
	cmd.Stdout = fakeStdout
	errRun := cmd.Run()
	if errRun != nil {
		dieOn(errRun, t)
	}
	dieOn(fakeStdout.Close(), t)
	newOutBytes, err := ioutil.ReadAll(r)
	dieOn(err, t)
	dieOn(r.Close(), t)
	return strings.TrimSuffix(string(newOutBytes), "\n")
}

func catchCat(t *testing.T, program string, input []string, stdin bool, nativeCat bool) string {
	var result string
	r, fakeStdout, err := os.Pipe()
	er, fakeStderr, _ := os.Pipe()
	dieOn(err, t)
	if stdin {
		cmd := exec.Command(program)
		stdin, err := cmd.StdinPipe()
		for i := range input {
			stdin.Write([]byte(input[i]))
		}
		stdin.Close()
		if err != nil {
			dieOn(err, t)
		}
		data, err := cmd.Output()
		if err != nil {
			dieOn(err, t)
		}
		result = string(data)
	} else {
		if nativeCat {
			cmd := exec.Command(program)
			for i := range input {
				cmd.Args = append(cmd.Args, input[i])
			}
			cmd.Stdout = fakeStdout
			errRun := cmd.Run()
			if errRun != nil {
				dieOn(errRun, t)
			}
			dieOn(fakeStdout.Close(), t)
			newOutBytes, err := ioutil.ReadAll(r)
			dieOn(err, t)
			dieOn(r.Close(), t)
			result = string(newOutBytes)

			r, fakeStdout, err = os.Pipe()
			cmd = exec.Command("cat")
			for i := range input {
				cmd.Args = append(cmd.Args, input[i])
			}
			cmd.Stdout = fakeStdout
			cmd.Stderr = fakeStderr
			errRun = cmd.Run()
			if errRun != nil {
				dieOn(fakeStderr.Close(), t)
				newBytes, _ := ioutil.ReadAll(er)
				newOutBytes = newBytes
			} else {
				dieOn(fakeStdout.Close(), t)
				newOutBytes, err = ioutil.ReadAll(r)
			}
			dieOn(err, t)
			dieOn(r.Close(), t)

			result += "|" + string(newOutBytes)
		} else {
			cmd := exec.Command(program)
			for i := range input {
				cmd.Args = append(cmd.Args, input[i])
			}

			cmd.Stdout = fakeStdout
			errRun := cmd.Run()
			if errRun != nil {
				dieOn(errRun, t)
			}
			dieOn(fakeStdout.Close(), t)
			newOutBytes, err := ioutil.ReadAll(r)
			dieOn(err, t)
			dieOn(r.Close(), t)
			result = string(newOutBytes)
		}
	}
	return result
}

func catchZtail(t *testing.T, program string, input []string) string {
	var result string
	r, fakeStdout, err := os.Pipe()
	er, fakeStderr, _ := os.Pipe()
	dieOn(err, t)
	cmd := exec.Command(program)
	for i := range input {
		cmd.Args = append(cmd.Args, input[i])
	}
	cmd.Stdout = fakeStdout
	errRun := cmd.Run()
	if errRun != nil {
		dieOn(errRun, t)
	}
	dieOn(fakeStdout.Close(), t)
	newOutBytes, err := ioutil.ReadAll(r)
	dieOn(err, t)
	dieOn(r.Close(), t)
	result = string(newOutBytes)

	r, fakeStdout, err = os.Pipe()
	dieOn(err, t)
	cmd = exec.Command("tail")
	for i := range input {
		cmd.Args = append(cmd.Args, input[i])
	}
	cmd.Stdout = fakeStdout
	cmd.Stderr = fakeStderr
	errRun = cmd.Run()
	var newBytes []byte
	if errRun != nil {
		dieOn(fakeStderr.Close(), t)
		newBytes, _ = ioutil.ReadAll(er)
	}
	dieOn(fakeStdout.Close(), t)
	newOutBytes, err = ioutil.ReadAll(r)
	dieOn(err, t)
	dieOn(r.Close(), t)
	result += "|" + string(newOutBytes) + string(newBytes)

	return result
}

func catchBTreeApplyByLevel(t *testing.T, runnable func(root *student.TreeNode, f func(...interface{}) (int, error)), root *student.TreeNode, f2 func(a ...interface{}) (int, error)) string {
	realStdout := os.Stdout
	defer func() { os.Stdout = realStdout }()
	r, fakeStdout, err := os.Pipe()
	dieOn(err, t)
	os.Stdout = fakeStdout
	runnable(root, f2)
	dieOn(fakeStdout.Close(), t)
	newOutBytes, err := ioutil.ReadAll(r)
	dieOn(err, t)
	dieOn(r.Close(), t)
	return strings.TrimSuffix(string(newOutBytes), "\n")
}

// TestAtoi function
func TestAtoi(task *testing.T) {
	var pass bool = true
	testing := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"-", 0},
		{"--123", 0},
		{"1", 1},
		{"-3", -3},
		{"8292", 8292},
		{"9223372036854775807", 9223372036854775807},
		{"-9223372036854775808", -9223372036854775808},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		result := student.Atoi(value.input)
		countPass++
		if result != value.expected {
			errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
			pass = false
			countPass--
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestRecursivePower(task *testing.T) {
	var pass bool = true
	testing := []struct {
		input    []int
		expected int
	}{
		{[]int{-7, -2}, 0},
		{[]int{-8, -7}, 0},
		{[]int{4, 8}, 65536},
		{[]int{1, 3}, 1},
		{[]int{-1, 1}, -1},
		{[]int{-6, 5}, -7776},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		result := student.RecursivePower(value.input[0], value.input[1])
		countPass++
		if result != value.expected {
			errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
			pass = false
			countPass--
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestPrintCombN(task *testing.T) {
	var pass bool = true
	testing := []struct {
		input    int
		expected string
	}{
		{1, "0, 1, 2, 3, 4, 5, 6, 7, 8, 9"},
		{2, "01, 02, 03, 04, 05, 06, 07, 08, 09, 12, 13, 14, 15, 16, 17, 18, 19, 23, 24, 25, 26, 27, 28, 29, 34, 35, 36, 37, 38, 39, 45, 46, 47, 48, 49, 56, 57, 58, 59, 67, 68, 69, 78, 79, 89"},
		{3, "012, 013, 014, 015, 016, 017, 018, 019, 023, 024, 025, 026, 027, 028, 029, 034, 035, 036, 037, 038, 039, 045, 046, 047, 048, 049, 056, 057, 058, 059, 067, 068, 069, 078, 079, 089, 123, 124, 125, 126, 127, 128, 129, 134, 135, 136, 137, 138, 139, 145, 146, 147, 148, 149, 156, 157, 158, 159, 167, 168, 169, 178, 179, 189, 234, 235, 236, 237, 238, 239, 245, 246, 247, 248, 249, 256, 257, 258, 259, 267, 268, 269, 278, 279, 289, 345, 346, 347, 348, 349, 356, 357, 358, 359, 367, 368, 369, 378, 379, 389, 456, 457, 458, 459, 467, 468, 469, 478, 479, 489, 567, 568, 569, 578, 579, 589, 678, 679, 689, 789"},
		{4, "0123, 0124, 0125, 0126, 0127, 0128, 0129, 0134, 0135, 0136, 0137, 0138, 0139, 0145, 0146, 0147, 0148, 0149, 0156, 0157, 0158, 0159, 0167, 0168, 0169, 0178, 0179, 0189, 0234, 0235, 0236, 0237, 0238, 0239, 0245, 0246, 0247, 0248, 0249, 0256, 0257, 0258, 0259, 0267, 0268, 0269, 0278, 0279, 0289, 0345, 0346, 0347, 0348, 0349, 0356, 0357, 0358, 0359, 0367, 0368, 0369, 0378, 0379, 0389, 0456, 0457, 0458, 0459, 0467, 0468, 0469, 0478, 0479, 0489, 0567, 0568, 0569, 0578, 0579, 0589, 0678, 0679, 0689, 0789, 1234, 1235, 1236, 1237, 1238, 1239, 1245, 1246, 1247, 1248, 1249, 1256, 1257, 1258, 1259, 1267, 1268, 1269, 1278, 1279, 1289, 1345, 1346, 1347, 1348, 1349, 1356, 1357, 1358, 1359, 1367, 1368, 1369, 1378, 1379, 1389, 1456, 1457, 1458, 1459, 1467, 1468, 1469, 1478, 1479, 1489, 1567, 1568, 1569, 1578, 1579, 1589, 1678, 1679, 1689, 1789, 2345, 2346, 2347, 2348, 2349, 2356, 2357, 2358, 2359, 2367, 2368, 2369, 2378, 2379, 2389, 2456, 2457, 2458, 2459, 2467, 2468, 2469, 2478, 2479, 2489, 2567, 2568, 2569, 2578, 2579, 2589, 2678, 2679, 2689, 2789, 3456, 3457, 3458, 3459, 3467, 3468, 3469, 3478, 3479, 3489, 3567, 3568, 3569, 3578, 3579, 3589, 3678, 3679, 3689, 3789, 4567, 4568, 4569, 4578, 4579, 4589, 4678, 4679, 4689, 4789, 5678, 5679, 5689, 5789, 6789"},
		{5, "01234, 01235, 01236, 01237, 01238, 01239, 01245, 01246, 01247, 01248, 01249, 01256, 01257, 01258, 01259, 01267, 01268, 01269, 01278, 01279, 01289, 01345, 01346, 01347, 01348, 01349, 01356, 01357, 01358, 01359, 01367, 01368, 01369, 01378, 01379, 01389, 01456, 01457, 01458, 01459, 01467, 01468, 01469, 01478, 01479, 01489, 01567, 01568, 01569, 01578, 01579, 01589, 01678, 01679, 01689, 01789, 02345, 02346, 02347, 02348, 02349, 02356, 02357, 02358, 02359, 02367, 02368, 02369, 02378, 02379, 02389, 02456, 02457, 02458, 02459, 02467, 02468, 02469, 02478, 02479, 02489, 02567, 02568, 02569, 02578, 02579, 02589, 02678, 02679, 02689, 02789, 03456, 03457, 03458, 03459, 03467, 03468, 03469, 03478, 03479, 03489, 03567, 03568, 03569, 03578, 03579, 03589, 03678, 03679, 03689, 03789, 04567, 04568, 04569, 04578, 04579, 04589, 04678, 04679, 04689, 04789, 05678, 05679, 05689, 05789, 06789, 12345, 12346, 12347, 12348, 12349, 12356, 12357, 12358, 12359, 12367, 12368, 12369, 12378, 12379, 12389, 12456, 12457, 12458, 12459, 12467, 12468, 12469, 12478, 12479, 12489, 12567, 12568, 12569, 12578, 12579, 12589, 12678, 12679, 12689, 12789, 13456, 13457, 13458, 13459, 13467, 13468, 13469, 13478, 13479, 13489, 13567, 13568, 13569, 13578, 13579, 13589, 13678, 13679, 13689, 13789, 14567, 14568, 14569, 14578, 14579, 14589, 14678, 14679, 14689, 14789, 15678, 15679, 15689, 15789, 16789, 23456, 23457, 23458, 23459, 23467, 23468, 23469, 23478, 23479, 23489, 23567, 23568, 23569, 23578, 23579, 23589, 23678, 23679, 23689, 23789, 24567, 24568, 24569, 24578, 24579, 24589, 24678, 24679, 24689, 24789, 25678, 25679, 25689, 25789, 26789, 34567, 34568, 34569, 34578, 34579, 34589, 34678, 34679, 34689, 34789, 35678, 35679, 35689, 35789, 36789, 45678, 45679, 45689, 45789, 46789, 56789"},
		{6, "012345, 012346, 012347, 012348, 012349, 012356, 012357, 012358, 012359, 012367, 012368, 012369, 012378, 012379, 012389, 012456, 012457, 012458, 012459, 012467, 012468, 012469, 012478, 012479, 012489, 012567, 012568, 012569, 012578, 012579, 012589, 012678, 012679, 012689, 012789, 013456, 013457, 013458, 013459, 013467, 013468, 013469, 013478, 013479, 013489, 013567, 013568, 013569, 013578, 013579, 013589, 013678, 013679, 013689, 013789, 014567, 014568, 014569, 014578, 014579, 014589, 014678, 014679, 014689, 014789, 015678, 015679, 015689, 015789, 016789, 023456, 023457, 023458, 023459, 023467, 023468, 023469, 023478, 023479, 023489, 023567, 023568, 023569, 023578, 023579, 023589, 023678, 023679, 023689, 023789, 024567, 024568, 024569, 024578, 024579, 024589, 024678, 024679, 024689, 024789, 025678, 025679, 025689, 025789, 026789, 034567, 034568, 034569, 034578, 034579, 034589, 034678, 034679, 034689, 034789, 035678, 035679, 035689, 035789, 036789, 045678, 045679, 045689, 045789, 046789, 056789, 123456, 123457, 123458, 123459, 123467, 123468, 123469, 123478, 123479, 123489, 123567, 123568, 123569, 123578, 123579, 123589, 123678, 123679, 123689, 123789, 124567, 124568, 124569, 124578, 124579, 124589, 124678, 124679, 124689, 124789, 125678, 125679, 125689, 125789, 126789, 134567, 134568, 134569, 134578, 134579, 134589, 134678, 134679, 134689, 134789, 135678, 135679, 135689, 135789, 136789, 145678, 145679, 145689, 145789, 146789, 156789, 234567, 234568, 234569, 234578, 234579, 234589, 234678, 234679, 234689, 234789, 235678, 235679, 235689, 235789, 236789, 245678, 245679, 245689, 245789, 246789, 256789, 345678, 345679, 345689, 345789, 346789, 356789, 456789"},
		{7, "0123456, 0123457, 0123458, 0123459, 0123467, 0123468, 0123469, 0123478, 0123479, 0123489, 0123567, 0123568, 0123569, 0123578, 0123579, 0123589, 0123678, 0123679, 0123689, 0123789, 0124567, 0124568, 0124569, 0124578, 0124579, 0124589, 0124678, 0124679, 0124689, 0124789, 0125678, 0125679, 0125689, 0125789, 0126789, 0134567, 0134568, 0134569, 0134578, 0134579, 0134589, 0134678, 0134679, 0134689, 0134789, 0135678, 0135679, 0135689, 0135789, 0136789, 0145678, 0145679, 0145689, 0145789, 0146789, 0156789, 0234567, 0234568, 0234569, 0234578, 0234579, 0234589, 0234678, 0234679, 0234689, 0234789, 0235678, 0235679, 0235689, 0235789, 0236789, 0245678, 0245679, 0245689, 0245789, 0246789, 0256789, 0345678, 0345679, 0345689, 0345789, 0346789, 0356789, 0456789, 1234567, 1234568, 1234569, 1234578, 1234579, 1234589, 1234678, 1234679, 1234689, 1234789, 1235678, 1235679, 1235689, 1235789, 1236789, 1245678, 1245679, 1245689, 1245789, 1246789, 1256789, 1345678, 1345679, 1345689, 1345789, 1346789, 1356789, 1456789, 2345678, 2345679, 2345689, 2345789, 2346789, 2356789, 2456789, 3456789"},
		{8, "01234567, 01234568, 01234569, 01234578, 01234579, 01234589, 01234678, 01234679, 01234689, 01234789, 01235678, 01235679, 01235689, 01235789, 01236789, 01245678, 01245679, 01245689, 01245789, 01246789, 01256789, 01345678, 01345679, 01345689, 01345789, 01346789, 01356789, 01456789, 02345678, 02345679, 02345689, 02345789, 02346789, 02356789, 02456789, 03456789, 12345678, 12345679, 12345689, 12345789, 12346789, 12356789, 12456789, 13456789, 23456789"},
		{9, "012345678, 012345679, 012345689, 012345789, 012346789, 012356789, 012456789, 013456789, 023456789, 123456789"},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		result := catchPrintCombN(task, student.PrintCombN, value.input)
		countPass++
		if result != value.expected {
			errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
			pass = false
			countPass--
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestPrintNbrBase(task *testing.T) {
	var pass bool = true
	testing := []struct {
		input    int
		base     string
		expected string
	}{
		{919617, "01", "11100000100001000001"},
		{753639, "CHOUMIisDAcat!", "HIDAHI"},
		{-174336, "CHOUMIisDAcat!", "-MssiD"},
		{-661165, "1", "NV"},
		{-861737, "Zone01Zone01", "NV"},
		{125, "0123456789ABCDEF", "7D"},
		{-125, "choumi", "-uoi"},
		{125, "-ab", "NV"},
		{-9223372036854775808, "0123456789", "-9223372036854775808"},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		result := catchPrintNbrBase(task, student.PrintNbrBase, value.input, value.base)
		countPass++
		if result != value.expected {
			errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
			pass = false
			countPass--
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestDoop(task *testing.T) {
	var pass bool = true
	testing := []struct {
		a, op, b string
		expected string
	}{
		{"861", "+", "870", "1731"},
		{"861", "-", "870", "-9"},
		{"861", "*", "870", "749070"},
		{"861", "%", "870", "861"},
		{"hello", "+", "1", "0"},
		{"1", "p", "1", "0"},
		{"1", "/", "0", "No division by 0"},
		{"1", "%", "0", "No modulo by 0"}, //"No remainder of division by 0"},
		{"1", "*", "1", "1"},
		{"9223372036854775807", "+", "1", "0"},
		{"9223372036854775809", "-", "3", "0"},
		{"9223372036854775807", "*", "3", "0"},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		result := catchDoop(task, "../doop/doop", value.a, value.op, value.b)
		countPass++
		resArray := strings.Split(value.expected, "|")
		if len(resArray) > 1 {
			if result != resArray[0] && result != resArray[1] {
				errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
				pass = false
				countPass--
			}
		} else {
			if result != resArray[0] {
				errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
				pass = false
				countPass--
			}
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestAtoibase(task *testing.T) {
	var pass bool = true
	testing := []struct {
		input    string
		base     string
		expected int
	}{
		{"bcbbbbaab", "abc", 12016},
		{"0001", "01", 1},
		{"00", "01", 0},
		{"saDt!I!sI", "CHOUMIisDAcat!", 11557277891},
		{"AAho?Ao", "WhoAmI?", 406772},
		{"thisinputshouldnotmatter", "abca", 0},
		{"125", "0123456789", 125},
		{"uoi", "choumi", 125},
		{"bbbbbab", "-ab", 0},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		result := student.AtoiBase(value.input, value.base)
		countPass++
		if result != value.expected {
			errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
			pass = false
			countPass--
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestSplitWhiteSpaces(task *testing.T) {
	var pass bool = true
	testing := []struct {
		input    string
		expected []string
	}{
		{"The earliest foundations of what would become computer science predate the invention of the modern digital computer", []string{"The", "earliest", "foundations", "of", "what", "would", "become", "computer", "science", "predate", "the", "invention", "of", "the", "modern", "digital", "computer"}},
		{"Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,", []string{"Machines", "for", "calculating", "fixed", "numerical", "tasks", "such", "as", "the", "abacus", "have", "existed", "since", "antiquity,"}},
		{"aiding in computations such as multiplication and division .", []string{"aiding", "in", "computations", "such", "as", "multiplication", "and", "division", "."}},
		{"Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment.", []string{"Algorithms", "for", "performing", "computations", "have", "existed", "since", "antiquity,", "even", "before", "the", "development", "of", "sophisticated", "computing", "equipment."}},
		{"Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]", []string{"Wilhelm", "Schickard", "designed", "and", "constructed", "the", "first", "working", "mechanical", "calculator", "in", "1623.[4]"}},
		{"In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,", []string{"In", "1673,", "Gottfried", "Leibniz", "demonstrated", "a", "digital", "mechanical", "calculator,"}},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		result := student.SplitWhiteSpaces(value.input)
		countPass++
		if len(result) == len(value.expected) {
			for i := range result {
				if result[i] != value.expected[i] {
					errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
					pass = false
					countPass--
				}
			}
		} else {
			errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
			pass = false
			countPass--
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestSplit(task *testing.T) {
	var pass bool = true
	testing := []struct {
		input    string
		charset  string
		expected []string
	}{
		{"|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,", "|=choumi=|", []string{"", "which", "itself", "used", "cards", "and", "a", "central", "computing", "unit.", "When", "the", "machine", "was", "finished,"}},
		{"!==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator,", "!==!", []string{"", "which", "was", "making", "all", "kinds", "of", "punched", "card", "equipment", "and", "was", "also", "in", "the", "calculator", "business[10]", "to", "develop", "his", "giant", "programmable", "calculator,"}},
		{"AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator,", "AFJ", []string{"", "Charles", "Babbage", "started", "the", "design", "of", "the", "first", "automatic", "mechanical", "calculator,"}},
		{"<<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer,", "<<==123==>>", []string{"", "In", "1820,", "Thomas", "de", "Colmar", "launched", "the", "mechanical", "calculator", "industry[note", "1]", "when", "he", "released", "his", "simplified", "arithmometer,"}},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		result := student.Split(value.input, value.charset)
		countPass++
		if len(result) == len(value.expected) {
			for i := range result {
				if result[i] != value.expected[i] {
					errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
					pass = false
					countPass--
				}
			}
		} else {
			errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
			pass = false
			countPass--
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestConvertBase(task *testing.T) {
	var pass bool = true
	testing := []struct {
		input            string
		basefrom, baseto string
		expected         string
	}{
		{"4506C", "0123456789ABCDEF", "choumi", "hccocimc"},
		{"babcbaaaaabcb", "abc", "0123456789ABCDEF", "9B611"},
		{"256850", "0123456789", "01", "111110101101010010"},
		{"uuhoumo", "choumi", "Zone01", "eeone0n"},
		{"683241", "0123456789", "0123456789", "683241"},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		result := student.ConvertBase(value.input, value.basefrom, value.baseto)
		countPass++
		for i := range result {
			if result[i] != value.expected[i] {
				errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
				pass = false
				countPass--
			}
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestRotateVowels(task *testing.T) {
	var pass bool = true
	testing := []struct {
		input    []string
		expected string
	}{
		{[]string{"Hello World"}, "Hollo Werld"},
		{[]string{"HEllO World", "problem solved"}, "Hello Werld problom sOlvEd"},
		{[]string{"str", "shh", "psst"}, "str shh psst"},
		{[]string{"happy thoughts", "good luck"}, "huppy thooghts guod lack"}, //"happy thoughts good luck"},
		{[]string{"al's elEphAnt is overly underweight!"}, "il's elephunt es ovirly AndErweaght!"},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		result := catchRotateVowels(task, "../rotatevowels/rotatevowels", value.input)
		countPass++
		if result != value.expected {
			errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
			pass = false
			countPass--
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestAdvancedSortWordArr(task *testing.T) {
	var pass bool = true
	testing := []struct {
		input    []string
		f        func(a, b string) int
		expected []string
	}{
		{[]string{"The", "earliest", "computing", "device", "undoubtedly", "consisted", "of", "the", "five", "fingers", "of", "each", "hand"}, strings.Compare, []string{"The", "computing", "consisted", "device", "each", "earliest", "fingers", "five", "hand", "of", "of", "the", "undoubtedly"}},
		{[]string{"The", "word", "digital", "comesfrom", "\"digits\"", "or", "fingers"}, strings.Compare, []string{"\"digits\"", "The", "comesfrom", "digital", "fingers", "or", "word"}},
		{[]string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}, strings.Compare, []string{"1", "2", "3", "A", "B", "C", "a", "b", "c"}},
		{[]string{"The", "computing", "consisted", "device", "each", "earliest", "fingers", "five", "hand", "of", "of", "the", "undoubtedly"}, func(a, b string) int { return strings.Compare(b, a) }, []string{"undoubtedly", "the", "of", "of", "hand", "five", "fingers", "earliest", "each", "device", "consisted", "computing", "The"}},
		{[]string{"\"digits\"", "The", "comesfrom", "digital", "fingers", "or", "word"}, func(a, b string) int { return strings.Compare(b, a) }, []string{"word", "or", "fingers", "digital", "comesfrom", "The", "\"digits\""}},
		{[]string{"1", "2", "3", "A", "B", "C", "a", "b", "c"}, func(a, b string) int { return strings.Compare(b, a) }, []string{"c", "b", "a", "C", "B", "A", "3", "2", "1"}},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		student.AdvancedSortWordArr(value.input, value.f)
		countPass++
		for i := range value.input {
			if value.input[i] != value.expected[i] {
				errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, value.input, value.expected))
				pass = false
				countPass--
			}
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestCat(task *testing.T) {
	var pass bool = true
	timeout := time.After(3 * time.Second)
	done := make(chan bool)
	testing := []struct {
		input            []string
		expected         string
		stdin, nativeCat bool
	}{
		{[]string{"Hello\n", "jaflsdfj\n", "Computer science (sometimes called computation science or computing science\n"}, "Hello\njaflsdfj\nComputer science (sometimes called computation science or computing science\n", true, false},
		{[]string{"../cat/quest8.txt"}, "\"Programming is a skill best acquired by pratice and example rather than from books\" by Alan Turing\n", false, false},
		{[]string{"../cat/quest8T.txt"}, "\"Alan Mathison Turing was an English mathematician, computer scientist, logician, cryptanalyst. Turing was highly influential in the development of theoretical computer science, providing a formalisation of the concepts of algorithm and computation with the Turing machine, which can be considered a model of a general-purpose computer. Turing is widely considered to be the father of theoretical computer science and artificial intelligence.\"\n", false, false},
		{[]string{"../cat/quest8T.txt", "../cat/quest8.txt"}, "\"Alan Mathison Turing was an English mathematician, computer scientist, logician, cryptanalyst. Turing was highly influential in the development of theoretical computer science, providing a formalisation of the concepts of algorithm and computation with the Turing machine, which can be considered a model of a general-purpose computer. Turing is widely considered to be the father of theoretical computer science and artificial intelligence.\"\n\"Programming is a skill best acquired by pratice and example rather than from books\" by Alan Turing\n", false, false},
		{[]string{"../cat/quest8.txt"}, "\"Programming is a skill best acquired by pratice and example rather than from books\" by Alan Turing\n|\"Programming is a skill best acquired by pratice and example rather than from books\" by Alan Turing\n", false, true},
		{[]string{"abcdefghijklmnopqrstqvwxyz"}, "open abcdefghijklmnopqrstqvwxyz: no such file or directory\n|cat: abcdefghijklmnopqrstqvwxyz: No such file or directory\n||cat: abcdefghijklmnopqrstqvwxyz: No such file or directory\n|cat: abcdefghijklmnopqrstqvwxyz: No such file or directory\n", false, true},
	}
	var countPass int
	var allCases = len(testing)
	go func() {
		for _, value := range testing {
			result := catchCat(task, "../cat/cat", value.input, value.stdin, value.nativeCat)
			countPass++
			resArray := strings.Split(value.expected, "||")
			if len(resArray) > 1 {
				if value.stdin {
					if result != resArray[0] && result != resArray[1] {
						errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
						pass = false
						countPass--
					}
				} else {
					if value.nativeCat {
						if result != resArray[0] && result != resArray[1] {
							errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
							pass = false
							countPass--
						}
					} else {
						if result != resArray[0] && result != resArray[1] {
							errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
							pass = false
							countPass--
						}
					}
				}
			} else {
				if value.stdin {
					if result != resArray[0] {
						errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
						pass = false
						countPass--
					}
				} else {
					if value.nativeCat {
						if result != resArray[0] {
							errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
							pass = false
							countPass--
						}
					} else {
						if result != resArray[0] {
							errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
							pass = false
							countPass--
						}
					}
				}
			}
		}
		done <- true
	}()
	select {
	case <-timeout:
		errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (All Test Cases timeout)"))
		pass = false
	case <-done:
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestZtail(task *testing.T) {
	var pass bool = true
	testing := []struct {
		input    []string
		expected string
	}{
		{[]string{"-c", "20", "../ztail/1.txt"}, "strates the problem:|strates the problem:"},
		{[]string{"../ztail/1.txt", "-c", "23"}, "monstrates the problem:|monstrates the problem:"},
		{[]string{"-c", "jelrjq", "15"}, "../ztail/ztail: invalid number of bytes: ‘jelrjq’\n|tail: invalid number of bytes: ‘jelrjq’\n||tail: invalid number of bytes: ‘jelrjq’\n|tail: invalid number of bytes: ‘jelrjq’\n"},
		{[]string{"-c", "11", "../ztail/1.txt", "jfdklsa"}, "==> ../ztail/1.txt <==\nhe problem:../ztail/ztail: cannot open 'jfdklsa' for reading: No such file or directory\n|==> ../ztail/1.txt <==\nhe problem:tail: cannot open 'jfdklsa' for reading: No such file or directory\n||==> ../ztail/1.txt <==\nhe problem:tail: cannot open 'jfdklsa' for reading: No such file or directory\n|==> ../ztail/1.txt <==\nhe problem:tail: cannot open 'jfdklsa' for reading: No such file or directory\n"},
		{[]string{"11", "-c", "../ztail/1.txt"}, "../ztail/ztail: invalid number of bytes: ‘../ztail/1.txt’\n|tail: invalid number of bytes: ‘../ztail/1.txt’\n||tail: invalid number of bytes: ‘../ztail/1.txt’\n|tail: invalid number of bytes: ‘../ztail/1.txt’\n"},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		result := catchZtail(task, "../ztail/ztail", value.input)
		countPass++
		resArray := strings.Split(value.expected, "||")
		if len(resArray) > 1 {
			if result != resArray[0] && result != resArray[1] {
				errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
				pass = false
				countPass--
			}
		} else {
			if result != resArray[0] {
				errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
				pass = false
				countPass--
			}
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestActiveBits(task *testing.T) {
	var pass bool = true
	testing := []struct {
		input    int
		expected uint
	}{
		{15, 4},
		{17, 2},
		{4, 1},
		{11, 3},
		{9, 2},
		{12, 2},
		{2, 1},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		result := student.ActiveBits(value.input)
		countPass++
		if result != value.expected {
			errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
			pass = false
			countPass--
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestSortListInsert(task *testing.T) {
	var pass bool = true
	testing := []struct {
		input    *student.NodeI
		dataref  int
		expected *student.NodeI
	}{
		{&student.NodeI{Data: 0}, 39, listPushBack(&student.NodeI{Data: 0}, 39)},
		{listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 0}, 1), 2), 3), 4), 24), 25), 54), 33, listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 0}, 1), 2), 3), 4), 24), 25), 33), 54)},
		{listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 0}, 2), 18), 33), 37), 37), 39), 52), 53), 57), 53, listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 0}, 2), 18), 33), 37), 37), 39), 52), 53), 53), 57)},
		{listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 0}, 5), 18), 24), 28), 35), 42), 45), 52), 52, listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 0}, 5), 18), 24), 28), 35), 42), 45), 52), 52)},
		{listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 0}, 12), 20), 23), 23), 24), 30), 41), 53), 57), 59), 38, listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 0}, 12), 20), 23), 23), 24), 30), 38), 41), 53), 57), 59)},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		result := student.SortListInsert(value.input, value.dataref)
		countPass++
		var currentResult = result
		var currentExpected = value.expected
		for currentExpected != nil {
			if currentResult == nil || currentResult.Data != currentExpected.Data {
				errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, PrintList(value.input), PrintList(value.expected)))
				pass = false
				countPass--
				break
			}
			currentExpected = currentExpected.Next
			currentResult = currentResult.Next
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestSortedListMerge(task *testing.T) {
	var pass bool = true
	testing := []struct {
		input1, input2 *student.NodeI
		expected       *student.NodeI
	}{
		{nil, nil, nil},
		{nil, listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 2}, 2), 4), 9), 12), 12), 19), 20), listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 2}, 2), 4), 9), 12), 12), 19), 20)},
		{listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 4}, 4), 6), 9), 13), 18), 20), 20), nil, listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 4}, 4), 6), 9), 13), 18), 20), 20)},
		{listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 0}, 7), 39), 92), 97), 93), 91), 28), 64), listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 80}, 23), 27), 30), 85), 81), 75), 70), listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 0}, 7), 23), 27), 30), 39), 70), 75), 80), 81), 85), 92), 97), 93), 91), 28), 64)},
		//{listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 0}, 7), 39), 92), 97), 93), 91), 28), 64), listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 80}, 23), 27), 30), 85), 81), 75), 70), listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 0}, 7), 23), 27), 28), 30), 39), 64), 70), 75), 80), 81), 85), 91), 92), 93), 97)},
		{listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 0}, 2), 11), 30), 54), 56), 70), 79), 99), listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 23}, 28), 38), 67), 67), 79), 95), 97), listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 0}, 2), 11), 23), 28), 30), 38), 54), 56), 67), 67), 70), 79), 79), 95), 97), 99)},
		{listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 0}, 3), 8), 8), 13), 19), 34), 38), 46), listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 7}, 39), 45), 53), 59), 70), 76), 79), listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(listPushBack(&student.NodeI{Data: 0}, 3), 7), 8), 8), 13), 19), 34), 38), 39), 45), 46), 53), 59), 70), 76), 79)},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		result := student.SortedListMerge(value.input1, value.input2)
		countPass++
		var currentResult = result
		var currentExpected = value.expected
		for currentExpected != nil {
			if currentResult == nil || currentResult.Data != currentExpected.Data {
				errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, PrintList(result), PrintList(value.expected)))
				pass = false
				countPass--
				break
			}
			currentExpected = currentExpected.Next
			currentResult = currentResult.Next
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestListRemoveIf(task *testing.T) {
	var pass bool = true
	testing := []struct {
		input    *student.List
		dataref  interface{}
		expected *student.List
	}{
		{&student.List{}, 1, &student.List{}},
		{&student.List{}, 96, &student.List{}},
		{ListPushBack(ListPushBack(ListPushBack(ListPushBack(ListPushBack(ListPushBack(ListPushBack(ListPushBack(ListPushBack(&student.List{}, 98), 98), 33), 34), 33), 34), 33), 89), 33), 34, ListPushBack(ListPushBack(ListPushBack(ListPushBack(ListPushBack(ListPushBack(ListPushBack(&student.List{}, 98), 98), 33), 33), 33), 89), 33)},
		{ListPushBack(ListPushBack(ListPushBack(ListPushBack(ListPushBack(&student.List{}, 79), 74), 99), 79), 7), 99, ListPushBack(ListPushBack(ListPushBack(ListPushBack(&student.List{}, 79), 74), 79), 7)},
		{ListPushBack(ListPushBack(ListPushBack(ListPushBack(ListPushBack(ListPushBack(ListPushBack(ListPushBack(&student.List{}, 56), 93), 68), 56), 87), 68), 56), 68), 68, ListPushBack(ListPushBack(ListPushBack(ListPushBack(ListPushBack(&student.List{}, 56), 93), 56), 87), 56)},
		{ListPushBack(ListPushBack(ListPushBack(&student.List{}, "mvkUxbqhQve4l"), "4Zc4t hnf SQ"), "q2If E8BPuX"), "4Zc4t hnf SQ", ListPushBack(ListPushBack(&student.List{}, "mvkUxbqhQve4l"), "q2If E8BPuX")},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		student.ListRemoveIf(value.input, value.dataref)
		countPass++
		var currentResult = value.input.Head
		var currentExpected = value.expected.Head
		for currentExpected != nil {
			if currentResult == nil || currentResult.Data != currentExpected.Data {
				errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, PrintList2(value.input), PrintList2(value.expected)))
				pass = false
				countPass--
				break
			}
			currentExpected = currentExpected.Next
			currentResult = currentResult.Next
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestBTreeTransplant(task *testing.T) {
	var pass bool = true
	tstRoot := &student.TreeNode{Data: "33"}
	tstRoot.Right = &student.TreeNode{Data: "5"}
	tstRoot.Right.Parent = tstRoot
	tstRoot.Left = &student.TreeNode{Data: "55"}
	tstRoot.Left.Parent = tstRoot
	tstRoot.Right.Right = &student.TreeNode{Data: "52"}
	tstRoot.Right.Right.Parent = tstRoot.Right
	tstRoot.Left.Right = &student.TreeNode{Data: "60"}
	tstRoot.Left.Right.Parent = tstRoot.Left
	tstRoot.Left.Left = &student.TreeNode{Data: "33"}
	tstRoot.Left.Left.Parent = tstRoot.Left
	tstRoot.Left.Left.Left = &student.TreeNode{Data: "12"}
	tstRoot.Left.Left.Left.Parent = tstRoot.Left.Left
	tstRoot.Left.Left.Left.Right = &student.TreeNode{Data: "15"}
	tstRoot.Left.Left.Left.Right.Parent = tstRoot.Left.Left.Left

	tstRoot2 := &student.TreeNode{Data: "03"}
	tstRoot2.Right = &student.TreeNode{Data: "39"}
	tstRoot2.Right.Parent = tstRoot2
	tstRoot2.Right.Right = &student.TreeNode{Data: "99"}
	tstRoot2.Right.Right.Parent = tstRoot2.Right
	tstRoot2.Right.Left = &student.TreeNode{Data: "55"}
	tstRoot2.Right.Left.Parent = tstRoot2.Right
	tstRoot2.Right.Right.Left = &student.TreeNode{Data: "44"}
	tstRoot2.Right.Right.Left.Parent = tstRoot2.Right.Right
	tstRoot2.Right.Left.Left = &student.TreeNode{Data: "33"}
	tstRoot2.Right.Left.Left.Parent = tstRoot2.Right.Left
	tstRoot2.Right.Left.Right = &student.TreeNode{Data: "60"}
	tstRoot2.Right.Left.Right.Parent = tstRoot2.Right.Left
	tstRoot2.Right.Left.Left.Left = &student.TreeNode{Data: "12"}
	tstRoot2.Right.Left.Left.Left.Parent = tstRoot2.Right.Left.Left
	tstRoot2.Right.Left.Left.Left.Right = &student.TreeNode{Data: "15"}
	tstRoot2.Right.Left.Left.Left.Right.Parent = tstRoot2.Right.Left.Left.Left

	testing := []struct {
		root, rplc, expected *student.TreeNode
		node                 string
	}{
		{BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(&student.TreeNode{Data: "01"}, "07"), "05"), "12"), "02"), "10"), "03"), BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(&student.TreeNode{Data: "55"}, "33"), "60"), "12"), "15"), BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(&student.TreeNode{Data: "01"}, "07"), "05"), "55"), "02"), "33"), "60"), "03"), "12"), "15"), "12"},
		{BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(&student.TreeNode{Data: "33"}, "20"), "5"), "13"), "31"), "52"), "11"), BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(&student.TreeNode{Data: "55"}, "33"), "60"), "12"), "15"), tstRoot, "20"},
		{BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(&student.TreeNode{Data: "03"}, "39"), "11"), "99"), "14"), "44"), "11"), BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(&student.TreeNode{Data: "55"}, "33"), "60"), "12"), "15"), tstRoot2, "11"},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		result := student.BTreeTransplant(value.root, BTreeSearchItem(value.root, value.node), value.rplc)
		countPass++
		if !BTreeIsEqual(result, value.expected) {
			errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d)\ngot:\n%v\nwant:\n%v", countPass, printBTree(result), printBTree(value.expected)))
			pass = false
			countPass--
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestBTreeApplyByLevel(task *testing.T) {
	var pass bool = true
	tstRoot := &student.TreeNode{Data: "33"}
	tstRoot.Right = &student.TreeNode{Data: "5"}
	tstRoot.Right.Parent = tstRoot
	tstRoot.Left = &student.TreeNode{Data: "55"}
	tstRoot.Left.Parent = tstRoot
	tstRoot.Right.Right = &student.TreeNode{Data: "52"}
	tstRoot.Right.Right.Parent = tstRoot.Right
	tstRoot.Left.Right = &student.TreeNode{Data: "60"}
	tstRoot.Left.Right.Parent = tstRoot.Left
	tstRoot.Left.Left = &student.TreeNode{Data: "33"}
	tstRoot.Left.Left.Parent = tstRoot.Left
	tstRoot.Left.Left.Left = &student.TreeNode{Data: "12"}
	tstRoot.Left.Left.Left.Parent = tstRoot.Left.Left
	tstRoot.Left.Left.Left.Right = &student.TreeNode{Data: "15"}
	tstRoot.Left.Left.Left.Right.Parent = tstRoot.Left.Left.Left

	tstRoot2 := &student.TreeNode{Data: "03"}
	tstRoot2.Right = &student.TreeNode{Data: "39"}
	tstRoot2.Right.Parent = tstRoot2
	tstRoot2.Right.Right = &student.TreeNode{Data: "99"}
	tstRoot2.Right.Right.Parent = tstRoot2.Right
	tstRoot2.Right.Left = &student.TreeNode{Data: "55"}
	tstRoot2.Right.Left.Parent = tstRoot2.Right
	tstRoot2.Right.Right.Left = &student.TreeNode{Data: "44"}
	tstRoot2.Right.Right.Left.Parent = tstRoot2.Right.Right
	tstRoot2.Right.Left.Left = &student.TreeNode{Data: "33"}
	tstRoot2.Right.Left.Left.Parent = tstRoot2.Right.Left
	tstRoot2.Right.Left.Right = &student.TreeNode{Data: "60"}
	tstRoot2.Right.Left.Right.Parent = tstRoot2.Right.Left
	tstRoot2.Right.Left.Left.Left = &student.TreeNode{Data: "12"}
	tstRoot2.Right.Left.Left.Left.Parent = tstRoot2.Right.Left.Left
	tstRoot2.Right.Left.Left.Left.Right = &student.TreeNode{Data: "15"}
	tstRoot2.Right.Left.Left.Left.Right.Parent = tstRoot2.Right.Left.Left.Left

	testing := []struct {
		root     *student.TreeNode
		f        func(a ...interface{}) (n int, err error)
		expected string
	}{
		{BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(&student.TreeNode{Data: "01"}, "07"), "05"), "12"), "02"), "10"), "03"), fmt.Println, "01\n07\n05\n12\n02\n10\n03"},
		{BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(&student.TreeNode{Data: "01"}, "07"), "05"), "12"), "02"), "10"), "03"), fmt.Print, "01070512021003"},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		result := catchBTreeApplyByLevel(task, student.BTreeApplyByLevel, value.root, value.f)
		countPass++
		if result != value.expected {
			errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d) got: %v, want: %v", countPass, result, value.expected))
			pass = false
			countPass--
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
}

func TestBTreeDeleteNode(task *testing.T) {
	var pass bool = true
	testing := []struct {
		root, expected *student.TreeNode
		node           string
	}{
		{BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(&student.TreeNode{Data: "01"}, "07"), "05"), "12"), "02"), "10"), "03"), BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(&student.TreeNode{Data: "01"}, "07"), "05"), "12"), "03"), "10"), "02"},
		{BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(&student.TreeNode{Data: "33"}, "20"), "5"), "13"), "31"), "52"), "11"), BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(&student.TreeNode{Data: "33"}, "31"), "5"), "13"), "52"), "11"), "20"},
		{BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(&student.TreeNode{Data: "03"}, "39"), "11"), "99"), "14"), "44"), "11"), BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(&student.TreeNode{Data: "39"}, "11"), "99"), "14"), "44"), "11"), "03"},
		{BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(&student.TreeNode{Data: "03"}, "01"), "03"), "94"), "19"), "111"), "24"), BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(BTreeInsertData(&student.TreeNode{Data: "03"}, "01"), "94"), "19"), "111"), "24"), "03"},
	}
	var countPass int
	var allCases = len(testing)
	for _, value := range testing {
		result := student.BTreeDeleteNode(value.root, BTreeSearchItem(value.root, value.node))
		countPass++
		if !BTreeIsEqual(result, value.expected) {
			errorsArray = append(errorsArray, fmt.Sprintf(strings.TrimPrefix(task.Name(), "Test")+" (Test Case #%d)\ngot:\n%v\nwant:\n%v", countPass, printBTree(result), printBTree(value.expected)))
			pass = false
			countPass--
		}
	}
	fmt.Printf("%-20s %s (%02d/%02d)\n", strings.TrimPrefix(task.Name(), "Test"), status(pass), countPass, allCases)
	fmt.Println()
}
func TestAll(task *testing.T) {
	var lenoferrorArray = len(errorsArray)

	if lenoferrorArray > 0 {
		fmt.Println()
		fmt.Println(ansi.ColorCode("yellow+bold") + "ERROR LIST:" + ansi.ColorCode("reset"))
		for i := range errorsArray {
			fmt.Println(errorsArray[i])
		}
		task.Fail()
	}
}
