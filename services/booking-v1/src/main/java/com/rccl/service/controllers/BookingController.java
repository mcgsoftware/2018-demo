package com.rccl.service.controllers;

import com.rccl.service.models.ErrorResponse;
import com.rccl.service.models.Reservation;
import com.rccl.service.models.ServiceResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.info.BuildProperties;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;

import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletResponse;

@RestController
@RequestMapping("/royal/api")
public class BookingController {
    private static Logger log = LoggerFactory.getLogger(BookingController.class);


    @Autowired
    BuildProperties buildProperties;

    @RequestMapping("/bookings/{vdsId}")
    public Reservation[] getShipInfo(@PathVariable String vdsId, HttpServletResponse response) {

        // Set ship code to 'N/A' so we know data came from this version of booking service

        Reservation r1 = new Reservation("bjm100", "NA1", "Alllure",
                "Oasis", "2018-12-30T00:00:00Z");
        Reservation r2 = new Reservation("bjm100", "NA2", "Symphony",
                "Oasis", "2019-03-14T00:00:00Z");
        Reservation[] reservations = {r1, r2};

        //String msg = "Feature not implemented to get info for: " + vdsId + " service version: " +
        //        buildProperties.getName() + " => " + buildProperties.getVersion();

        // Header hack for SPA demo app.
        response.setHeader("Access-Control-Allow-Origin", "*");

        return(reservations);
    }
}
