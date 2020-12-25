package twofive

const D = 20201227

func findLoopSize(publicKey, subjectNumber int) int {
	value := 1

	for ls := 1; ; ls++ {
		value = (value * subjectNumber) % D
		if value == publicKey {
			return ls
		}
	}
}

func findEncryptionKey(loopSize, subjectNumber int) int {
	value := 1
	for ls := 0; ls < loopSize; ls++ {
		value = (value * subjectNumber) % D
	}
	return value
}
