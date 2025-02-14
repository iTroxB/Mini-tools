#!/bin/bash

generate_rut() {
    first_two=$((RANDOM % 40 + 60))

    middle_six=$(printf "%06d" $((RANDOM % 1000000)))

    base_number="$first_two$middle_six"

    reverse_number=$(echo "$base_number" | rev)
    multiplier=(2 3 4 5 6 7)
    sum=0

    for ((i = 0; i < ${#reverse_number}; i++)); do
        digit=${reverse_number:$i:1}
        sum=$((sum + digit * multiplier[i % ${#multiplier[@]}]))
    done

    remainder=$((sum % 11))
    verifier=$((11 - remainder))

    if [ "$verifier" -eq 11 ]; then
        verifier=0
    elif [ "$verifier" -eq 10 ]; then
        verifier="K"
    fi

    echo "$base_number$verifier"
}

count=0
while [ $count -lt 50 ]; do
    echo "Creando RUTs chilenos $((count + 1))..."
    generate_rut
    ((count++))
done

echo
exit 0
