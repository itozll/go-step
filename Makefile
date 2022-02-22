TARGET := go-step

.PHONY: $(TARGET)

all: $(TARGET)

$(TARGET):
	go build -o $@ main.go