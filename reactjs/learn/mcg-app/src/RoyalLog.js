
export class Base {
    constructor(
        eventType,
        traceId,
        userId,
        datetime,
        feature
    ) {
        this.eventType = eventType
        this.traceId = traceId
        this.userId = userId
        this.datetime = datetime
        this.feature = feature
    }

    static factoryError(traceId, vdsId) {
        const base = new Base("Error", traceId, vdsId, 
          new Date().toISOString(), 'profile')
        return base
    }
}

export class AppInfo {
    constructor(
        app,
        platform,
        version,
        device,
        datacenter
    ) {
        this.app = app
        this.platform = platform
        this.version = version
        this.device = device
        this.datacenter = datacenter
    }

     static factory() {
         let ai = new AppInfo("ReactDemo", "web", "1.0", "chrome", "GCP")
         return ai
     }
}

export class ErrorInfo {
    constructor(
        errId,
        errMsg,
        blame,
        context,
        validations,
        stack,
        errRate,
        extErrId
    ) {
        this.errId = errId
        this.errMsg = errMsg
        this.blame = blame
        this.context = context
        this.validations = validations
        this.stack = stack
        this.errRate = errRate
        this.extErrId = extErrId
    }
}

export class ErrorEvent {
    constructor(
        event,    // Base 
        appInfo,  // AppInfo
        error     // ErrorInfo
    ) {
        this.event = event
        this.appInfo = appInfo
        this.error = error 
    }
}