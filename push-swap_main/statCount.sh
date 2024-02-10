rm statRes.txt

for i in {1..10000}
do
  go run . >> statRes.txt
done