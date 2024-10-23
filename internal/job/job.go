package job

// Flavor is a definition of VM resource.
type Flavor struct {
	Cores int    // Minimum number of CPU cores required.
	Ram   int    // Minimum amount of memory required in GB.
	Disk  int    // Minimum amount of memory required in GB.
	ID    string // Flavor UUID if available.
}

type Image struct {
	ID   string
	Name string
}

type Job struct {
	Flavor *Flavor
	Image  *Image
}
