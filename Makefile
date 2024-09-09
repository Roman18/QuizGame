
all:
	go build -ldflags "-w -s"


PHONY:
	clean

clean:
	rm QuizGame
