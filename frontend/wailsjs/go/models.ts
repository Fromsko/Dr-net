export namespace core {
	
	export class NetworkPair {
	    content: {[key: string]: []};
	
	    static createFrom(source: any = {}) {
	        return new NetworkPair(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.content = source["content"];
	    }
	}

}

export namespace main {
	
	export class ApiIOResult {
	    flag: boolean;
	    data: string;
	
	    static createFrom(source: any = {}) {
	        return new ApiIOResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.flag = source["flag"];
	        this.data = source["data"];
	    }
	}

}

