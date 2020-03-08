<template>
	<div>
		<v-btn
			@click="tooglePlay"
			class="md-icon-button md-theme-footer md-primary"
		>
			<v-icon>{{ playerIcon }}</v-icon>
		</v-btn>
		<v-btn class="md-icon-button md-theme-footer md-primary">
			<v-icon>volume_down</v-icon>
		</v-btn>
		<v-btn class="md-icon-button md-theme-footer md-primary">
			<v-icon>volume_up</v-icon>
		</v-btn>
		<audio id="player" :src="''"></audio>
	</div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
	data() {
		return {
			player: {},
			playerStatus: "paused",
			playerIcon: "play_arrow",
		};
	},

	props: {
		source: {
			type: String,
			default: "http://dream01.gr:8000/stream",
		},
	},

	created() {},

	computed: {
		...mapGetters(["getPlayerStatus", "getPlayerSource"]),
	},

	methods: {
		...mapActions(["setPlayerStatus"]),

		tooglePlay() {
			var player = document.getElementById("player");

			if (this.getPlayerStatus == "paused") {
				//play
				this.playerIcon = "pause";

				player.pause();
				player.src = this.getPlayerSource;
				player.load();
				player.play();

				this.setPlayerStatus("playing");
			} else {
				//pause
				this.playerIcon = "play_arrow";
				player.pause();
				this.setPlayerStatus("paused");
			}
		},
	},

	watch: {
		getPlayerSource: val => {
			var player = document.getElementById("player");

			var url = new URL(val);

			player.pause();
			player.src = url.href;
			player.load();
			player.play().then(
				() => {},
				() => {
					player.src = url.href + ";";
					player.load();
					player.play();
				}
			);
		},
	},
};
</script>

<style lang="scss"></style>
