module github.com/julibarri/go-uc3m-training

go 1.21.1

// replace asks to replace the mentioned package with the path that you
// mentioned, so it won't further look packages elsewhere and
// would look inside that's handlers package located there itself
replace github.com/julibarri/go-uc3m-training/hello_world => ./hello_world


require github.com/julibarri/go-uc3m-training/hello_world v0.0.0-20240513101340-64197acbb13e
