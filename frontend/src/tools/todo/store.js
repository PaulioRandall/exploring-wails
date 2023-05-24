import { writable } from 'svelte/store'

export const tasks = writable([])

let idPool = 0
export const genId = () => {
	idPool++
	return idPool
}
