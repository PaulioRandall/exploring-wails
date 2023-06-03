export namespace database {
	export class Task {
		id: number
		text: string

		static createFrom(source: any = {}) {
			return new Task(source)
		}

		constructor(source: any = {}) {
			if ('string' === typeof source) source = JSON.parse(source)
			this.id = source['id']
			this.text = source['text']
		}
	}
}

export namespace files {
	export class ReadOnlyFile {
		name: string
		path: string
		absPath: string
		isDir: boolean

		static createFrom(source: any = {}) {
			return new ReadOnlyFile(source)
		}

		constructor(source: any = {}) {
			if ('string' === typeof source) source = JSON.parse(source)
			this.name = source['name']
			this.path = source['path']
			this.absPath = source['absPath']
			this.isDir = source['isDir']
		}
	}
}
