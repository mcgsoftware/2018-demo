package com.rccl.service;

import io.javalin.Javalin;
import com.rccl.service.models.ResponseMessage;

import javax.servlet.http.HttpServletResponse;

public class HelloWorld {

	public static void main(String[] args) {
		Javalin app = Javalin.create().start(7000);
		app.get("/", ctx -> ctx.result("Hello World"));

		app.get("/royal/api/login/health", ctx -> ctx.json(new ResponseMessage("Up")) );

		app.post("/royal/api/login", ctx -> {

			String u = ctx.formParam("username");
			String p = ctx.formParam("password");
			if ((u == null || p == null)) {
				ResponseMessage err = new ResponseMessage("Missing html form username or password");
				ctx.status(HttpServletResponse.SC_NOT_FOUND);
				ctx.json(err);
			} else if (("sri".equals(u)) && ("brian".equals(p))) {
				ctx.json(new ResponseMessage("Login successful."));
			} else {
				ResponseMessage err = new ResponseMessage("Invalid username or password");
				ctx.status(HttpServletResponse.SC_NOT_FOUND);
				ctx.json(err);
			}
		});

	}

}
