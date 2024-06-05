function printFlag() {
    const flag = "43 57 56 56 33 44 36 6b 21 5d 4f 69 20 65 62 4f 56 5c 24 57 6d";
    const hexArray = flag.split(" ");

    let asciiString = hexArray.map(hex => {
        return String.fromCharCode(parseInt(hex, 16) + 16);
    });
    asciiString = asciiString.join("");

    console.log(asciiString);
}
