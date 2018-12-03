package com.rccl.service.models;

public class Reservation implements ServiceResponse {
    private final String vdsId;
    private final String shipCode;
    private final String shipName;
    private final String shipClass;
    private final String sailDate;

    public Reservation(String vdsId, String shipCode, String shipName, String shipClass, String sailDate) {
        this.vdsId = vdsId;
        this.shipCode = shipCode;
        this.shipName = shipName;
        this.shipClass = shipClass;
        this.sailDate = sailDate;
    }


    public String getVdsId() {
        return vdsId;
    }

    public String getShipCode() {
        return shipCode;
    }

    public String getShipName() {
        return shipName;
    }

    public String getShipClass() {
        return shipClass;
    }

    public String getSailDate() {
        return sailDate;
    }
}
