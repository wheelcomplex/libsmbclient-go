#include<stdio.h>
#include "_cgo_export.h"


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

ssize_t my_smbc_read(SMBCCTX *c, SMBCFILE *file, void *buf, size_t count) {
  smbc_read_fn fn = smbc_getFunctionRead(c);
  return fn(c, file, buf, count);
}

void my_smbc_close(SMBCCTX *c, SMBCFILE *f) {
  smbc_close_fn fn = smbc_getFunctionClose(c);
  fn(c, f);
}


extern void 
GoAuthCallbackHelper(char *, char *, char *, int, char *, int, char *, int);
void my_smbc_auth_callback(SMBCCTX *context,
	       const char *server_name, const char *share_name,
	       char *domain_out, int domainmaxlen,
	       char *username_out, int unmaxlen,
	       char *password_out, int pwmaxlen) {
   GoAuthCallbackHelper((char*)server_name, (char*)share_name, domain_out,  domainmaxlen, username_out, unmaxlen, password_out, pwmaxlen);
   //FIXME: check {domain,u,pw}maxlen
}

void my_smbc_init_auth_callback(SMBCCTX *ctx)
{
   smbc_setFunctionAuthDataWithContext(ctx, my_smbc_auth_callback);
}