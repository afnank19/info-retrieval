CXX = g++
CXXFLAGS = -Iinclude
SRC = ./src/*.cpp
OUT = main

run: $(SRC)
	$(CXX) $(CXXFLAGS) $(SRC) -o $(OUT)
	./$(OUT)

build: $(SRC)
	$(CXX) $(CXXFLAGS) $(SRC) -o $(OUT)

run-build:
	./$(OUT)

clean:
	rm ./$(OUT)