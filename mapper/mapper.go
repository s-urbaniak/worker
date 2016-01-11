package mapper

// the request communication channel to the mapper
var req = make(chan interface{})

type wordMsg struct {
	word string
	cnt chan int
}

type (
	addWordMsg wordMsg
	cntWordMsg wordMsg
)

func newWordMsg(w string) wordMsg {
	return wordMsg{w, make(chan int)}
}

// AddWord adds a word to the mapper and returns the current count
func AddWord(w string) int {
	msg := newWordMsg(w)
	req <- addWordMsg(msg)
	return <-msg.cnt
}

// WordCnt returns the word count for the given word
func WordCnt(w string) int {
	msg := newWordMsg(w)
	req <- cntWordMsg(msg)
	return <-msg.cnt
}

func mapper() {
	// a map mapping words to counts
	words := map[string]int{}

	for msg := range req {
		switch v := msg.(type) {
		case addWordMsg:
			words[string(v.word)] += 1
			cnt, _ := words[string(v.word)]
			v.cnt <- cnt

		case cntWordMsg:
			cnt, _ := words[string(v.word)]
			v.cnt <- cnt
		}
	}
}

func init() {
	go mapper()
}
