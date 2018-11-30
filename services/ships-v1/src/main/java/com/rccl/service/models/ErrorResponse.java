package com.rccl.service.models;

public class ErrorResponse implements ServiceResponse {
    private final long errno;
    private final String errmsg;

    public ErrorResponse(long errno, String errmsg) {
        this.errno = errno;
        this.errmsg = errmsg;
    }

    public long getErrno() { return errno; }
    public String getErrmsg() { return errmsg; }

}