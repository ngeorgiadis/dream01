<template>
	<div>
		{{ radios }}
	</div>
</template>

<script>
import axios from "axios";
export default {
	data() {
		return {
			radios: [],
		};
	},
	created() {
		axios.get("/fav").then(res => {
			this.radios = res.data;
		});
	},
};
</script>


<style lang="scss">
</style>
