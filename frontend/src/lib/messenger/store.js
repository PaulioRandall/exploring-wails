import { writable } from 'svelte/store'

export const message = writable(null)

const newMessage = (type, text) => {
	return {
		type,
		text,
	}
}

export const newSuccessMsg = (text) => newMessage('success', text)
export const newErrorMsg = (text) => newMessage('error', text)

export const msgTypes = {
	success: {
		bgColor: 'green',
	},
	error: {
		bgColor: 'indianred',
	},
}
