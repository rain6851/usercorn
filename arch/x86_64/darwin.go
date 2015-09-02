package x86_64

import (
	"../../models"
	uc "github.com/lunixbochs/unicorn"
)

var darwinSyscalls = map[int]string{
	1:   "exit",
	2:   "fork",
	3:   "read",
	4:   "write",
	5:   "open",
	6:   "close",
	7:   "wait4",
	9:   "link",
	10:  "unlink",
	73:  "munmap",
	197: "mmap",
	199: "lseek",
}

func DarwinSyscall(u models.Usercorn) {
	rax, _ := u.RegRead(uc.UC_X86_REG_RAX)
	AbiSyscall(u, darwinSyscalls, rax-0x2000000)
}

func DarwinInterrupt(u models.Usercorn, intno uint32) {
	if intno == 0x80 {
		DarwinSyscall(u)
	}
}

func init() {
	Arch.RegisterOS(&models.OS{Name: "darwin", Syscall: DarwinSyscall, Interrupt: DarwinInterrupt})
}
