<template>
	<div class="home md-layout">
		<md-card v-for="r in radios" :key="r.id" class="md-layout-item  ">
			<md-card-header>
				<md-card-header-text>
					<div class="md-title">Media card</div>
					<div class="md-subhead">Normal size</div>
				</md-card-header-text>

				<md-card-media>
					<img src="../assets/logo.png" alt="People" />
				</md-card-media>
			</md-card-header>

			<md-card-actions>
				<md-button>Action</md-button>
				<md-button>Action</md-button>
			</md-card-actions>
		</md-card>
	</div>
</template>

<script>
// @ is an alias to /src

export default {
	name: "home",
	components: {},
	data() {
		return {
			radios: [
				{ id: "1", name: "radio1", src: "123" },
				{ id: "2", name: "radio1", src: "123" },
				{ id: "3", name: "radio1", src: "123" },
				{ id: "4", name: "radio1", src: "123" },
				{ id: "5", name: "radio1", src: "123" },
				{ id: "6", name: "radio1", src: "123" },
				{ id: "7", name: "radio1", src: "123" },
				{ id: "8", name: "radio1", src: "123" },
				{ id: "9", name: "radio1", src: "123" },
				{ id: "10", name: "radio1", src: "123" },
				{ id: "11", name: "radio1", src: "123" }
			]
		};
	}
};
</script>

<style lang="scss" scoped>
.md-card {
	min-width: 250px;
	max-width: 250px;
	margin-bottom: 15px;
}
</style>
