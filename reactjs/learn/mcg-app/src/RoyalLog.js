
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

     
}

export class ServiceMetricInfo {
    constructor(
        service,
        operation,
        method,
        latency_ms,
        tags, 
        context
    ) {
        this.service = service
        this.operation = operation
        this.method = method
        this.latency_ms = latency_ms
        this.tags = tags
        this.context = context
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

export class ServiceMetricEvent {
    constructor(
        event,      // Base 
        app,        // AppInfo
        metrics     // ServiceMetricInfo
    ) {
        this.event = event
        this.app = app
        this.metrics = metrics  
    }
}