# Enigma and Lorenz
This is a collection of utilities for encrypting messages using virtual Enigma and Lorenz machines.

## Binaries
Pre-built binaries for Windows, Linux, and MacOS (ARM) are available in the `out` folder (64 bit only).

## Enigma

To get a list of possible commands run `enigma` with a `-h` flag:
```
Usage of enigma:
-c string
Center rotor number (I-VIII), position (1-26), and ring setting (0-25) (default "II 1 0")
-f string
Fourth rotor (beta|gamma), position (1-26), and ring setting (0-25) [optional]
-l string
Left rotor number (I-VIII), position (1-26), and ring setting (0-25) (default "I 1 0")
-m string
The message to be encrypted/decrypted
-plugs string
Plug mappings in the form of 'A:B C:D'[optional], position, and ring setting
-r string
Right rotor number (I-VIII), position (1-26), and ring setting (0-25) (default "III 1 0")
-ukw string
Reflector to use (A|B|C|a|b) (default "B")
```

### Example Input
#### Using default rotor settings
```sh
$ enigma -m "hello world"
ILBDAAMTAZ
$ enigma -m "ILBDAAMTAZ"
HELLOWORLD
```

#### Using custom rotor settings
```sh
$ enigma -m "hello world" -l "I 4 13" -c "IV 13 24" -r "II 12 23" -ukw "C"
JBVTUFTEYI
$ enigma -m "JBVTUFTEYI" -l "I 4 13" -c "IV 13 24" -r "II 12 23" -ukw "C"
HELLOWORLD
```

## Lorenz

To get a list of possible commands run `enigma` with a `-h` flag:
```
Usage of lorenz:
  -psi string
        The rotor setting for the Psi wheels (default "0 0 0 0 0")
  -chi string
        The rotor setting for the Chi wheels (0-max) (default "0 0 0 0 0")
  -d    Whether you are seeking to decrypt a message (0-max)
  -m string
        The message to be encrypted/decrypted
  -mot string
        The rotor setting for the Motor wheels (0-max) (default "0 0")

```

### Example Input
#### Using default rotor settings
```sh
$ lorenz -m "hello world"
K_BB|DWTTP^
$ lorenz -d -m "K_BB|DWTTP^" # Note the -d for decrypt
HELLO WORLD
```

#### Using custom rotor settings
```sh
$ lorenz -m "hello world" -psi "1 3 14 5 6" -chi "23 14 5 6" -mot "30 17"
B RLIGMLLUN
$ lorenz -d -m "B RLIGMLLUN" -psi "1 3 14 5 6" -chi "23 14 5 6" -mot "30 17"
HELLO WORLD
```