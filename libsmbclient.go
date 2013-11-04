package libsmbclient

/*
#cgo LDFLAGS: -lsmbclient

#include <libsmbclient.h>
#include <stdlib.h>
#include <unistd.h>

SMBCFILE* my_smbc_opendir(SMBCCTX *c, const char *fname) {
  smbc_opendir_fn fn = smbc_getFunctionOpendir(c);
  return fn(c, fname);
}

int my_smbc_closedir(SMBCCTX *c, SMBCFILE *dir) {
  smbc_closedir_fn fn = smbc_getFunctionClosedir(c);
  return fn(c, dir);
}

struct smbc_dirent* my_smbc_readdir(SMBCCTX *c, SMBCFILE *dir) {
  smbc_readdir_fn fn = smbc_getFunctionReaddir(c);
  return fn(c, dir);
}

SMBCFILE* my_smbc_open(SMBCCTX *c, const char *fname, int flags, mode_t mode) {
  smbc_open_fn fn = smbc_getFunctionOpen(c);
  return fn(c, fname, flags, mode);
}

void my_smbc_close(SMBCCTX *c, SMBCFILE *f) {
  smbc_close_fn fn = smbc_getFunctionClose(c);
  return fn(c, f);
}

*/
import (
	"C"
	"unsafe"
)

type Dirent struct {
	/** Type of entity.
	    SMBC_WORKGROUP=1,
	    SMBC_SERVER=2, 
	    SMBC_FILE_SHARE=3,
	    SMBC_PRINTER_SHARE=4,
	    SMBC_COMMS_SHARE=5,
	    SMBC_IPC_SHARE=6,
	    SMBC_DIR=7,
	    SMBC_FILE=8,
	    SMBC_LINK=9,*/ 
	Type int
	Comment string
	Name string
}

// Global init
func Init(debug int) error {
	_, err := C.smbc_init(nil, C.int(debug))
	return err
}

type File struct {
	smbcfile *C.SMBCFILE
}

// client interface
type Client struct {
	ctx *C.SMBCCTX
}

// debug stuff
func (c *Client) GetDebug() int {
	return int(C.smbc_getDebug(c.ctx));
}

func (c *Client) SetDebug(level int)  {
	C.smbc_setDebug(c.ctx, C.int(level));
}

func (c *Client) Init() {
	c.ctx = C.smbc_new_context();
	C.smbc_init_context(c.ctx)
}

func  (c *Client) Opendir(durl string) (File, error) {
	d, err := C.my_smbc_opendir(c.ctx, C.CString(durl))
	return File{d}, err
}

func (c *Client) Closedir(dir File) error {
	_, err := C.my_smbc_closedir(c.ctx, dir.smbcfile)
	return err
}

func (c *Client) Readdir(dir File) (Dirent, error) {
	c_dirent, err := C.my_smbc_readdir(c.ctx, dir.smbcfile)
	dirent := Dirent{Type: int(c_dirent.smbc_type),
		         Comment: C.GoString(c_dirent.comment),
		         Name: C.GoString(&c_dirent.name[0])}
	return dirent, err
}




// FIXME: mode is actually "mode_t mode"
func (c *Client) Open(furl string, flags int, mode int) (int, error) {
	cs := C.CString(furl)
	sf, err := C.my_smbc_open(c.ctx, cs, C.int(flags), C.mode_t(mode))
	return 	File{smbcfile: sf}, err
}

func (c *Client) Close(f File) {
	C.my_smbc_close(c.ctx, f.smbcfile)
}


