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
		{{ currentStationInfo }}
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
		...mapGetters([
			"getPlayerStatus",
			"getPlayerSource",
			"getCurrentStation",
		]),

		currentStationInfo() {
			let c = this.getCurrentStation;
			if (c === undefined || c.name === undefined) {
				return "";
			}

			return `${c.name} (${c.genre}) ${c.lc} - ${c.ct}`;
		},
	},

	methods: {
		...mapActions(["setPlayerStatus"]),

		tooglePlay() {
			if (this.getPlayerStatus == "paused") {
				this.play();
			} else {
				this.pause();
			}
		},

		/*var player = document.getElementById("player");

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
			} */

		play(source) {
			try {
				let player = document.getElementById("player");
				let url = "";
				if (!source) {
					url = new URL(source);
				} else {
					url = new URL(this.getPlayerSource);
				}

				player.pause();
				player.src = url.href;
				player.load();

				player.play().then(
					() => {
						this.setPlayerStatus("playing");
					},
					() => {
						player.src = url.href + ";";
						player.load();
						player.play().then(
							() => {
								this.setPlayerStatus("playing");
							},
							() => {
								this.setPlayerStatus("paused");
							}
						);
					}
				);
				this.playerIcon = "pause";
			} catch (e) {
				alert("cannot play..." + e);
				this.setPlayerStatus("paused");
				this.playerIcon = "play_arrow";
			}
		},

		pause() {
			let player = document.getElementById("player");
			this.playerIcon = "play_arrow";
			player.pause();
			this.setPlayerStatus("paused");
		},
	},

	watch: {
		getPlayerSource: function(val) {
			this.play(val);
			// try {
			// 	var player = document.getElementById("player");
			// 	var url = new URL(val);

			// 	player.pause();
			// 	player.src = url.href;
			// 	player.load();

			// 	player.play().then(
			// 		() => {
			// 			this.setPlayerStatus("playing");
			// 		},
			// 		() => {
			// 			player.src = url.href + ";";
			// 			player.load();
			// 			player.play().then(
			// 				() => {
			// 					this.setPlayerStatus("playing");
			// 				},
			// 				() => {
			// 					this.setPlayerStatus("paused");
			// 				}
			// 			);
			// 		}
			// 	);
			// } catch (e) {
			// 	alert("cannot play..." + e);
			// 	this.setPlayerStatus("paused");
			// }
		},
	},
};
</script>

<style lang="scss"></style>
