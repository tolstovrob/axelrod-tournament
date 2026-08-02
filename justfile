@play:
  go run ./cmd/play/
 
@experiments *args:
  go run ./cmd/experiments/ {{args}}
 
@exp0:
  go run ./cmd/experiments/ -e 0
 
@exp1:
  go run ./cmd/experiments/ -e 1
 
@exp2:
  go run ./cmd/experiments/ -e 2
 
@exp2s:
  go run ./cmd/experiments/ -e 2s
 
@exp3:
  go run ./cmd/experiments/ -e 3
 
@exp4:
  go run ./cmd/experiments/ -e 4
 
@exp5:
  go run ./cmd/experiments/ -e 5
 
@exp7:
  go run ./cmd/experiments/ -e 7
 
@all:
  go run ./cmd/experiments/ -e all
 