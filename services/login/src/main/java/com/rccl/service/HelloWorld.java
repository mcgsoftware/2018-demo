package com.rccl.service;

import io.javalin.Javalin;

public class HelloWorld {

	public static void main(String[] args) {
		Javalin app = Javalin.create().start(7000);
		app.get("/", ctx -> ctx.result("Hello World"));
		app.get("/royal/api/login", ctx -> ctx.result("Hello Login 1"));
		app.get("/royal/api/login/foo", ctx -> ctx.result("Hello Login foo"));

	}

}
