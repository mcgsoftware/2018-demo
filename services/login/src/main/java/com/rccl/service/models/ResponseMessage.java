package com.rccl.service.models;

public class ResponseMessage {
    private final String msg;

    public ResponseMessage(String msg) {
        this.msg = msg;
    }

    public String getMsg() { return msg; }
}