package com.rccl.service.controllers;

import com.rccl.service.models.ErrorResponse;
import com.rccl.service.models.ServiceResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.info.BuildProperties;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;

import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/royal/api")
public class ShipController {
    private static Logger log = LoggerFactory.getLogger(ShipController.class);


    @Autowired
    BuildProperties buildProperties;

    @RequestMapping("/ships/{shipCode}")
    public ServiceResponse getShipInfo(@PathVariable String shipCode) {
        String msg = "Feature not implemented to get info for: " + shipCode + " service version: " +
                buildProperties.getName() + " => " + buildProperties.getVersion();
        return(new ErrorResponse(1200, msg));
    }
}
