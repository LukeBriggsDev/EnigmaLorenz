package main

import (
	"EnigmaLorenz/pkg/lorenz"
	"EnigmaLorenz/pkg/util"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validateMessage(message string) (string, error) {
	message = strings.ToUpper(message)
	if !util.ValidChars(message, true) {
		return "", errors.New("invalid character found in text")
	}
	return message, nil
}

func validateChiPsiPositions(positions string, wheels [5]lorenz.Wheel) ([5]lorenz.Wheel, error) {
	splitPos := strings.Split(positions, " ")
	if len(splitPos) != 5 {
		return wheels, errors.New("invalid number of positions given")
	}

	for idx, pos := range splitPos {
		num, err := strconv.Atoi(pos)
		if err != nil {
			return wheels, errors.New("non numeric value given for rotor position")
		}

		if num < 0 || num > len(wheels[idx].GetPins()) {
			return wheels, errors.New("rotor position for Chi/Psi is not valid for that rotor")
		}

		wheels[idx].SetPos(byte(num))

	}

	return wheels, nil

}

func validateMotorPositions(positions string) ([2]lorenz.Wheel, error) {
	wheels := lorenz.NewWheelSet().Motor
	splitPos := strings.Split(positions, " ")
	if len(splitPos) != 2 {
		return wheels, errors.New("invalid number of positions given")
	}

	for idx, pos := range splitPos {
		num, err := strconv.Atoi(pos)
		if err != nil {
			return wheels, errors.New("non numeric value given for rotor position")
		}

		if num < 0 || num > len(wheels[idx].GetPins()) {
			return wheels, errors.New("rotor position for motor is not valid for that rotor")
		}

		wheels[idx].SetPos(byte(num))

	}

	return wheels, nil
}

func main() {
	messagePtr := flag.String("m", "", "The message to be encrypted/decrypted")
	chiPositionsPtr := flag.String("chi", "0 0 0 0 0", "The rotor setting for the Chi wheels (0-max)")
	mPositionsPtr := flag.String("mot", "0 0", "The rotor setting for the Motor wheels (0-max)")
	psiPositionsPtr := flag.String("Psi", "0 0 0 0 0", "The rotor setting for the Psi wheels")
	decryptPtr := flag.Bool("d", false, "Whether you are seeking to decrypt a message (0-max)")

	flag.Parse()

	message, err := validateMessage(*messagePtr)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Invalid characters in message, can only include A-Z, 0-9")
	}

	chiWheels, err := validateChiPsiPositions(*chiPositionsPtr, lorenz.NewWheelSet().Chi)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error for chi wheels: %s", err)
	}

	psiWheels, err := validateChiPsiPositions(*psiPositionsPtr, lorenz.NewWheelSet().Psi)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error for psi wheels: %s", err)
	}

	motorWheels, err := validateMotorPositions(*mPositionsPtr)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error for motor wheels: %s", err)
	}

	machine := lorenz.NewLorenz(chiWheels, motorWheels, psiWheels)

	alphabet := lorenz.NewITA2LSB()
	encoded, _ := alphabet.AsciiToITA2(message, *decryptPtr)
	encrypted := machine.Encrypt(encoded)
	decoded, _ := alphabet.ITA2ToAscii(encrypted, *decryptPtr)
	fmt.Printf("%s\n", decoded)

}
