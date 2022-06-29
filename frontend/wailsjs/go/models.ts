export namespace app {
	
	export class ParseResult {
	    scheme?: string;
	    opaque?: string;
	    host?: string;
	    path?: string;
	    raw_path?: string;
	    force_query?: boolean;
	    raw_query?: string;
	    fragment?: string;
	    raw_fragment?: string;
	
	    static createFrom(source: any = {}) {
	        return new ParseResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.scheme = source["scheme"];
	        this.opaque = source["opaque"];
	        this.host = source["host"];
	        this.path = source["path"];
	        this.raw_path = source["raw_path"];
	        this.force_query = source["force_query"];
	        this.raw_query = source["raw_query"];
	        this.fragment = source["fragment"];
	        this.raw_fragment = source["raw_fragment"];
	    }
	}

}

