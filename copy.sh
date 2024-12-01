
# Small util for copying the state of the day1 folder and replicating it across the other 24 directories.
# Copies / renames the directory, then replaces all references to "day1" with "dayX"
for dayNum in {2..25}; do
  rm -r "day$dayNum"
  cp -r day1 "day$dayNum"
  sed "s/day1/day$dayNum/g" "day1/main.go" > day"$dayNum"/main.go
done