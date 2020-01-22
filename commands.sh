# metronome

metronome play 120 4 4 -o stdout

metronome play 120 4 4 -o audio


# controller

controller server --storage-path ~/src/github.com/aklinkert/godays-lighting-talk/demo-storage --disable-controller

# defer

time go run code/counter.go

time go run code/counter.go -defer

# DMX timing

go run code/music_fps.go

go run code/channels_fps.go
