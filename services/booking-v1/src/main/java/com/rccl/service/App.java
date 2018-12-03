package com.rccl.service;


import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.info.BuildProperties;


import java.net.InetAddress;

@SpringBootApplication
public class App {

    public static String HOSTNAME = "unknown";


    private static final Logger LOGGER = LoggerFactory.getLogger(App.class);

    @Autowired
    BuildProperties buildProperties;



    // Run this via: http://localhost:8080/  per application.properties
    public static void main(String[] args) {
        try {
            InetAddress ip = InetAddress.getLocalHost();
            HOSTNAME = ip.getHostName();
        } catch (Throwable t) {
            t.printStackTrace();
            System.exit(1);
        }

        SpringApplication.run(App.class, args);
    }

}