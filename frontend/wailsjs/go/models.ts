export namespace app {
	
	
	
	export class Detail {
	    quality?: number;
	    popularity?: number;
	    maintenance?: number;
	
	    static createFrom(source: any = {}) {
	        return new Detail(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.quality = source["quality"];
	        this.popularity = source["popularity"];
	        this.maintenance = source["maintenance"];
	    }
	}
	export class Score {
	    final?: number;
	    // Go type: Detail
	    detail?: any;
	
	    static createFrom(source: any = {}) {
	        return new Score(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.final = source["final"];
	        this.detail = this.convertValues(source["detail"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Maintenance {
	    releasesFrequency?: number;
	    commitsFrequency?: number;
	    openIssues?: number;
	    issuesDistribution?: number;
	
	    static createFrom(source: any = {}) {
	        return new Maintenance(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.releasesFrequency = source["releasesFrequency"];
	        this.commitsFrequency = source["commitsFrequency"];
	        this.openIssues = source["openIssues"];
	        this.issuesDistribution = source["issuesDistribution"];
	    }
	}
	export class Popularity {
	    communityInterest?: number;
	    downloadsCount?: number;
	    downloadsAcceleration?: number;
	    dependentsCount?: number;
	
	    static createFrom(source: any = {}) {
	        return new Popularity(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.communityInterest = source["communityInterest"];
	        this.downloadsCount = source["downloadsCount"];
	        this.downloadsAcceleration = source["downloadsAcceleration"];
	        this.dependentsCount = source["dependentsCount"];
	    }
	}
	export class Quality {
	    carefulness?: number;
	    tests?: number;
	    health?: number;
	    branding?: number;
	
	    static createFrom(source: any = {}) {
	        return new Quality(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.carefulness = source["carefulness"];
	        this.tests = source["tests"];
	        this.health = source["health"];
	        this.branding = source["branding"];
	    }
	}
	export class Evaluation {
	    // Go type: Quality
	    quality?: any;
	    // Go type: Popularity
	    popularity?: any;
	    // Go type: Maintenance
	    maintenance?: any;
	
	    static createFrom(source: any = {}) {
	        return new Evaluation(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.quality = this.convertValues(source["quality"], null);
	        this.popularity = this.convertValues(source["popularity"], null);
	        this.maintenance = this.convertValues(source["maintenance"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Modifiers {
	    branch?: string;
	
	    static createFrom(source: any = {}) {
	        return new Modifiers(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.branch = source["branch"];
	    }
	}
	export class Info {
	    service?: string;
	    type?: string;
	    // Go type: Modifiers
	    modifiers?: any;
	
	    static createFrom(source: any = {}) {
	        return new Info(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.service = source["service"];
	        this.type = source["type"];
	        this.modifiers = this.convertValues(source["modifiers"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Urls {
	    original?: string;
	    service?: string;
	    shields?: string;
	    content?: string;
	
	    static createFrom(source: any = {}) {
	        return new Urls(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.original = source["original"];
	        this.service = source["service"];
	        this.shields = source["shields"];
	        this.content = source["content"];
	    }
	}
	export class Badges {
	    // Go type: Urls
	    urls?: any;
	    // Go type: Info
	    info?: any;
	
	    static createFrom(source: any = {}) {
	        return new Badges(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.urls = this.convertValues(source["urls"], null);
	        this.info = this.convertValues(source["info"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Files {
	    readmeSize?: number;
	    testsSize?: number;
	    hasNpmIgnore?: boolean;
	    hasChangelog?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Files(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.readmeSize = source["readmeSize"];
	        this.testsSize = source["testsSize"];
	        this.hasNpmIgnore = source["hasNpmIgnore"];
	        this.hasChangelog = source["hasChangelog"];
	    }
	}
	export class Source {
	    // Go type: Files
	    files?: any;
	    badges?: Badges[];
	    linters?: string[];
	    outdatedDependencies?: {[key: string]: string};
	
	    static createFrom(source: any = {}) {
	        return new Source(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.files = this.convertValues(source["files"], null);
	        this.badges = this.convertValues(source["badges"], Badges);
	        this.linters = source["linters"];
	        this.outdatedDependencies = source["outdatedDependencies"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Statuses {
	    context?: string;
	    state?: string;
	
	    static createFrom(source: any = {}) {
	        return new Statuses(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.context = source["context"];
	        this.state = source["state"];
	    }
	}
	export class Commits {
	    // Go type: time.Time
	    from?: any;
	    // Go type: time.Time
	    to?: any;
	    count?: number;
	
	    static createFrom(source: any = {}) {
	        return new Commits(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.from = this.convertValues(source["from"], null);
	        this.to = this.convertValues(source["to"], null);
	        this.count = source["count"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Contributors {
	    username?: string;
	    commitsCount?: number;
	
	    static createFrom(source: any = {}) {
	        return new Contributors(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.username = source["username"];
	        this.commitsCount = source["commitsCount"];
	    }
	}
	export class Issues {
	    count?: number;
	    openCount?: number;
	    isDisabled?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Issues(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.count = source["count"];
	        this.openCount = source["openCount"];
	        this.isDisabled = source["isDisabled"];
	    }
	}
	export class Github {
	    homepage?: string;
	    starsCount?: number;
	    forksCount?: number;
	    subscribersCount?: number;
	    // Go type: Issues
	    issues?: any;
	    contributors?: Contributors[];
	    commits?: Commits[];
	    statuses?: Statuses[];
	
	    static createFrom(source: any = {}) {
	        return new Github(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.homepage = source["homepage"];
	        this.starsCount = source["starsCount"];
	        this.forksCount = source["forksCount"];
	        this.subscribersCount = source["subscribersCount"];
	        this.issues = this.convertValues(source["issues"], null);
	        this.contributors = this.convertValues(source["contributors"], Contributors);
	        this.commits = this.convertValues(source["commits"], Commits);
	        this.statuses = this.convertValues(source["statuses"], Statuses);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Downloads {
	    // Go type: time.Time
	    from?: any;
	    // Go type: time.Time
	    to?: any;
	    count?: number;
	
	    static createFrom(source: any = {}) {
	        return new Downloads(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.from = this.convertValues(source["from"], null);
	        this.to = this.convertValues(source["to"], null);
	        this.count = source["count"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Npm {
	    downloads?: Downloads[];
	    starsCount?: number;
	
	    static createFrom(source: any = {}) {
	        return new Npm(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.downloads = this.convertValues(source["downloads"], Downloads);
	        this.starsCount = source["starsCount"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Releases {
	    // Go type: time.Time
	    from?: any;
	    // Go type: time.Time
	    to?: any;
	    count?: number;
	
	    static createFrom(source: any = {}) {
	        return new Releases(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.from = this.convertValues(source["from"], null);
	        this.to = this.convertValues(source["to"], null);
	        this.count = source["count"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Links {
	    npm?: string;
	    homepage?: string;
	    repository?: string;
	    bugs?: string;
	
	    static createFrom(source: any = {}) {
	        return new Links(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.npm = source["npm"];
	        this.homepage = source["homepage"];
	        this.repository = source["repository"];
	        this.bugs = source["bugs"];
	    }
	}
	export class Repository {
	    type?: string;
	    url?: string;
	    directory?: string;
	
	    static createFrom(source: any = {}) {
	        return new Repository(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.url = source["url"];
	        this.directory = source["directory"];
	    }
	}
	export class Maintainers {
	    username?: string;
	    email?: string;
	
	    static createFrom(source: any = {}) {
	        return new Maintainers(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.username = source["username"];
	        this.email = source["email"];
	    }
	}
	export class Publisher {
	    username?: string;
	    email?: string;
	
	    static createFrom(source: any = {}) {
	        return new Publisher(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.username = source["username"];
	        this.email = source["email"];
	    }
	}
	export class Author {
	    name?: string;
	
	    static createFrom(source: any = {}) {
	        return new Author(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	    }
	}
	export class Metadata {
	    name?: string;
	    scope?: string;
	    version?: string;
	    description?: string;
	    keywords?: string[];
	    // Go type: time.Time
	    date?: any;
	    // Go type: Author
	    author?: any;
	    // Go type: Publisher
	    publisher?: any;
	    maintainers?: Maintainers[];
	    // Go type: Repository
	    repository?: any;
	    // Go type: Links
	    links?: any;
	    license?: string;
	    dependencies?: {[key: string]: string};
	    devDependencies?: {[key: string]: string};
	    peerDependencies?: {[key: string]: string};
	    releases?: Releases[];
	    readme?: string;
	
	    static createFrom(source: any = {}) {
	        return new Metadata(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.scope = source["scope"];
	        this.version = source["version"];
	        this.description = source["description"];
	        this.keywords = source["keywords"];
	        this.date = this.convertValues(source["date"], null);
	        this.author = this.convertValues(source["author"], null);
	        this.publisher = this.convertValues(source["publisher"], null);
	        this.maintainers = this.convertValues(source["maintainers"], Maintainers);
	        this.repository = this.convertValues(source["repository"], null);
	        this.links = this.convertValues(source["links"], null);
	        this.license = source["license"];
	        this.dependencies = source["dependencies"];
	        this.devDependencies = source["devDependencies"];
	        this.peerDependencies = source["peerDependencies"];
	        this.releases = this.convertValues(source["releases"], Releases);
	        this.readme = source["readme"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Collected {
	    // Go type: Metadata
	    metadata?: any;
	    // Go type: Npm
	    npm?: any;
	    // Go type: Github
	    github?: any;
	    // Go type: Source
	    source?: any;
	
	    static createFrom(source: any = {}) {
	        return new Collected(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.metadata = this.convertValues(source["metadata"], null);
	        this.npm = this.convertValues(source["npm"], null);
	        this.github = this.convertValues(source["github"], null);
	        this.source = this.convertValues(source["source"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class GetPackageInsightsResponse {
	    // Go type: time.Time
	    analyzedAt?: any;
	    // Go type: Collected
	    collected?: any;
	    // Go type: Evaluation
	    evaluation?: any;
	    // Go type: Score
	    score?: any;
	
	    static createFrom(source: any = {}) {
	        return new GetPackageInsightsResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.analyzedAt = this.convertValues(source["analyzedAt"], null);
	        this.collected = this.convertValues(source["collected"], null);
	        this.evaluation = this.convertValues(source["evaluation"], null);
	        this.score = this.convertValues(source["score"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ParseUrlResponse {
	    scheme?: string;
	    opaque?: string;
	    host?: string;
	    path?: string;
	    raw_path?: string;
	    force_query?: boolean;
	    raw_query?: string;
	    fragment?: string;
	    raw_fragment?: string;
	    query?: {[key: string]: string[]};
	
	    static createFrom(source: any = {}) {
	        return new ParseUrlResponse(source);
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
	        this.query = source["query"];
	    }
	}

}

