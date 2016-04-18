#include "memory.h"

#include <fstream>
#include <iostream>
#include <stdint.h>

using namespace std;

void memory::loadFromFile(string fileName) {
    unsigned char bytes[4];

    uint32_t length = 0;

    FILE *fp=fopen(fileName.c_str(),"rb");
    while ( fread(bytes, 2, 1, fp) != 0) {
        length++;
    }
    fclose(fp);

    this->memory = new uint32_t[length];
    this->memoryLength = length;

    fp=fopen(fileName.c_str(),"rb");
    uint32_t index = 0;
    while ( fread(bytes, 2, 1, fp) != 0) {

        this->memory[index] = bytes[0] | (bytes[1]<<8);
        index++;
    }
    fclose(fp);
}

uint32_t memory::read() {
    uint32_t currentPointer = this->memoryPointer;
    this->memoryPointer++;

    return this->memory[currentPointer];
}

uint32_t memory::read(uint32_t pointer) {
    return this->memory[pointer];
}

void memory::write(uint32_t pointer, uint32_t value) {
    this->memory[pointer] = value;
}

uint32_t memory::getPointer() {
    return this->memoryPointer;
}

void memory::setPointer(uint32_t pointer) {
    this->memoryPointer = pointer;
}

bool memory::isEOM() {
    return this->memoryPointer >= this->memoryLength;
}