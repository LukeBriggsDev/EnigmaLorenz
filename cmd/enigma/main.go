package main

import (
	"EnigmaLorenz/pkg/util"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)
import "EnigmaLorenz/pkg/enigma"

// validateRotorInput takes the user's rotor parameter and returns the corresponding Rotor.
// An error is returned in cases where the input is not valid.
//
// Errors
//
// The returned error will not be nil if:
//	- There are not 3 arguments seperated by a space
//	- The first argument is not one of the valid rotors (I - VIII|beta|gamma|)
//	- The rotor setting is not between 1 and 26 inclusively
//	- The ring setting is not between 0 and 25 exclusively
//
func validateRotorInput(input string) (enigma.Rotor, error) {
	args := strings.Split(input, " ")
	rotorSet := enigma.GenerateRotors()
	rotor := rotorSet.I

	if len(args) != 3 {
		return rotor, errors.New("incorrect number of arguments")
	}

	// Validate rotor wheel
	switch args[0] {
	case "I":
		rotor = rotorSet.I
	case "II":
		rotor = rotorSet.II
	case "III":
		rotor = rotorSet.III
	case "IV":
		rotor = rotorSet.IV
	case "V":
		rotor = rotorSet.V
	case "VI":
		rotor = rotorSet.VI
	case "VII":
		rotor = rotorSet.VII
	case "VIII":
		rotor = rotorSet.VIII
	case "beta":
		rotor = rotorSet.Beta
	case "gamma":
		rotor = rotorSet.Gamma
	default:
		return rotor, errors.New("rotor selection is invalid")
	}

	// Validate rotor position
	pos, err := strconv.Atoi(args[1])
	if err != nil {
		return rotor, err
	}
	if pos < 1 || pos > 26 {
		return rotor, errors.New("rotor position not between 1 and 26")
	}

	rotor.SetShownPos(byte(pos))

	// Validate ring setting
	ring, err := strconv.Atoi(args[2])
	if err != nil {
		return rotor, err
	}
	if ring < 0 || ring > 25 {
		return rotor, errors.New("ring setting not between 0 and 25")
	}

	rotor.SetRingSetting(byte(ring))

	return rotor, nil

}

// validateReflectorInput takes the user's reflector parameter and returns the corresponding Rotor for that reflector.
// an error is returned in cases where the parameter is not valid
//
// Errors
//
// The returned error will not be nil if the rotor is not one of the specified reflectors (A, B, C, b, c)
func validateReflectorInput(input string) (enigma.Rotor, error) {
	rotorSet := enigma.GenerateRotors()
	switch input {
	case "A":
		return rotorSet.UKW_A, nil
	case "B":
		return rotorSet.UKW_B, nil
	case "C":
		return rotorSet.UKW_C, nil
	case "b":
		return rotorSet.UKW_b, nil
	case "c":
		return rotorSet.UKW_c, nil
	default:
		return rotorSet.UKW_B, errors.New("incorrect value for reflector")
	}
}

// validatePlugboardInput takes the user's plugboard parameter and returns the corresponding Plugboard.
// an error is returned in cases where the parameter is not valid
//
// Errors
//
// The returned error will not be nil if:
//	- There are not a character either side of the mapping
//	- The character is not between A-Z
// 	- There is more than one character either side of the mapping
//
func validatePlugboardInput(input string) (enigma.Plugboard, error) {

	plugboard := enigma.NewPlugboard()
	if len(input) == 0 {
		return plugboard, nil
	}

	mappings := strings.Split(input, " ")

	for _, mapping := range mappings {
		split := strings.Split(mapping, ":")
		if len(split) != 2 {
			return plugboard, errors.New("incorrect format for plugboard")
		}

		if !util.ValidChars(strings.ToUpper(split[0]), false) || !util.ValidChars(strings.ToUpper(split[1]), false) {
			return plugboard, errors.New("incorrect characters passed to plugboard")
		}

		if len(split[0]) != 1 || len(split[1]) != 1 {
			return plugboard, errors.New("invalid mapping")
		}

		plugboard.AddPlug(strings.ToUpper(split[0])[0], strings.ToUpper(split[1])[0])
	}

	return plugboard, nil
}

func main() {
	messagePtr := flag.String("m", "", "The message to be encrypted/decrypted")

	leftRotorPtr := flag.String("l", "I 1 0", "Left rotor number (I-VIII), position (1-26), and ring setting (0-25)")
	centerRotorPtr := flag.String("c", "II 1 0", "Center rotor number (I-VIII), position (1-26), and ring setting (0-25)")
	rightRotorPtr := flag.String("r", "III 1 0", "Right rotor number (I-VIII), position (1-26), and ring setting (0-25)")
	fourthRotorPtr := flag.String("f", "", "Fourth rotor (beta|gamma), position (1-26), and ring setting (0-25) [optional]")
	reflectorPtr := flag.String("ukw", "B", "Reflector to use (A|B|C|a|b)")
	plugsPtr := flag.String("plugs", "", "Plug mappings in the form of 'A:B C:D'[optional], position, and ring setting")

	flag.Parse()

	leftRotor, err := validateRotorInput(*leftRotorPtr)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error for left rotor: %s\n", err)
		os.Exit(1)
	}

	centerRotor, err := validateRotorInput(*centerRotorPtr)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error for center rotor: %s\n", err)
		os.Exit(1)
	}

	rightRotor, err := validateRotorInput(*rightRotorPtr)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error for right rotor: %s\n", err)
		os.Exit(1)
	}

	fourthRotor, err := validateRotorInput(*fourthRotorPtr)
	useFourthRotor := true
	if err != nil {
		if *fourthRotorPtr == "" {
			useFourthRotor = false
		} else {
			_, _ = fmt.Fprintf(os.Stderr, "Error for fourth rotor: %s\n", err)
			os.Exit(1)
		}
	}

	reflector, err := validateReflectorInput(*reflectorPtr)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error for reflector: %s\n", err)
		os.Exit(1)
	}

	plugs, err := validatePlugboardInput(*plugsPtr)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error for plugboard: %s\n", err)
		os.Exit(1)
	}

	machine := enigma.Enigma{
		LeftRotor:   leftRotor,
		CenterRotor: centerRotor,
		RightRotor:  rightRotor,
		FourthRotor: fourthRotor,
		Reflector:   reflector,
		Plugs:       plugs,
	}

	message := strings.Replace(strings.ToUpper(*messagePtr), " ", "", -1)

	if !util.ValidChars(message, false) {
		_, _ = fmt.Fprintf(os.Stderr, "Invalid characters in message: %s", message)
	}

	cipher, err := machine.Encrypt(message, useFourthRotor)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Encryption failed: %s", err)
	}

	fmt.Println(cipher)

}
