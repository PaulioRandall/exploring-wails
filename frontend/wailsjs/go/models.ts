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
