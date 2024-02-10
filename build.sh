# Compile les deux programmes dans le fichier principal

cd push-swap_main
go build -o ../push-swap
cd ../checker_main
go build -o ../checker
cd ..