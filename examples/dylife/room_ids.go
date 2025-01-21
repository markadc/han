package main

import (
	"fmt"
	"kss"
	"kss/reqs"
	"log"
)

func main() {
	client := reqs.Client{}
	api := "https://eos.douyin.com/life/api/live_screen/v4/replay/live_room_list"
	payload := kss.S{
		"user_id":            "1258293549605997",
		"begin_date":         "2024-12-23",
		"end_date":           "2025-01-21",
		"compare_begin_date": "2024-11-23",
		"compare_end_date":   "2024-12-22",
	}
	cookie := "UIFID_TEMP=ad7b1e526ce029fae15c4fe388b5a88a6cfc5b31e687fc2dca11932336017cb5716355992e9cd8df74978b7ec82322567b5afc5ce2a4c99d36cc8b94e1cc83a98dc3cb91d1cc096b25f19cfc8229d94c; bd_ticket_guard_client_web_domain=2; live_use_vvc=%22false%22; _bd_ticket_crypt_doamin=2; __security_server_data_status=1; store-region=cn-hb; store-region-src=uid; SEARCH_RESULT_LIST_TYPE=%22single%22; my_rd=2; SelfTabRedDotControl=%5B%7B%22id%22%3A%227380374513367189543%22%2C%22u%22%3A633%2C%22c%22%3A0%7D%2C%7B%22id%22%3A%226945392252973746184%22%2C%22u%22%3A158%2C%22c%22%3A0%7D%5D; passport_mfa_token=CjfhBAum5BPrr7%2BeNuRLChYyTDjrOvvxPjT5rNAbJqWKUcrKGj532ZvsN1ZwjhM9JNnE26j2CqybGkoKPHmDmv04Ac%2BNBx4GtkgiFy%2BWbIAisJgjDGMFI9duvh9jXeA2JY98qstJFpWWN6ushhYLM47g%2F8KaeMJ64RC%2B%2F%2BMNGPax0WwgAiIBA49cGRc%3D; d_ticket=945ab34fb972cfb15b2453ff0f69f2ae90505; login_time=1734431406707; _bd_ticket_crypt_cookie=c6767c3ec7163a71b42197ac015f40a8; is_staff_user=false; is_dash_user=0; passport_csrf_token=e6b323be411bf60a06ecf56b5e82841e; passport_csrf_token_default=e6b323be411bf60a06ecf56b5e82841e; __security_mc_1_s_sdk_crypt_sdk=68dbf446-46ee-ac20; __security_mc_1_s_sdk_cert_key=dbf4d9c6-4708-9e6c; __security_mc_1_s_sdk_sign_data_key_web_protect=066751cc-4d61-ad97; __security_mc_1_s_sdk_sign_data_key_sso=bf66b7ef-407a-ba3b; volume_info=%7B%22isMute%22%3Atrue%2C%22isUserMute%22%3Afalse%2C%22volume%22%3A0.43%7D; s_v_web_id=verify_m5uu1cdh_8TH7mVKC_bGNR_4bpi_9y8j_5HPuYNo9Ib1e; download_guide=%223%2F20250116%2F0%22; csrf_session_id=a1b8ff09a40a3eb0e14fa2586bd492f2; be-token=; be-token-scene=; passport_assist_user=Cjw8aYzIgRRdUHxr-YK2geNF15_8R2Sko4lV699l5tckx7Zc0r2jfMZ7rHWXT370ux_MEaDk-fHBXfEFSb0aSgo87OJg09QypRc20ODeA2a8HiC8DhuHLvKIL0Z7yys5zS-kEONEZBnZqkqwlC5Gm-dlqWX4WbRZYwv8oLKLEPr85g0Yia_WVCABIgEDfeP2TA%3D%3D; n_mh=vr2dnfNak7Pt-Bhzu_rNBe87CtXQ7wylgakLkCK1630; sid_guard=444a4b6ecf0fcd736ff7dc35656ba470%7C1737013323%7C5183999%7CMon%2C+17-Mar-2025+07%3A42%3A02+GMT; uid_tt=6c9e1a5dad980bae62535f0613f0179a; uid_tt_ss=6c9e1a5dad980bae62535f0613f0179a; sid_tt=444a4b6ecf0fcd736ff7dc35656ba470; sessionid=444a4b6ecf0fcd736ff7dc35656ba470; sessionid_ss=444a4b6ecf0fcd736ff7dc35656ba470; sid_ucp_v1=1.0.0-KDRlYmJjNWRiYmZkZDY0YzY1OWFkODBjZGIxYzgzNGM4NWFhNmNiZmMKHgjN_KjV8QIQy_CivAYY5Z8TIAww7fKo2AU4B0D0BxoCaGwiIDQ0NGE0YjZlY2YwZmNkNzM2ZmY3ZGMzNTY1NmJhNDcw; ssid_ucp_v1=1.0.0-KDRlYmJjNWRiYmZkZDY0YzY1OWFkODBjZGIxYzgzNGM4NWFhNmNiZmMKHgjN_KjV8QIQy_CivAYY5Z8TIAww7fKo2AU4B0D0BxoCaGwiIDQ0NGE0YjZlY2YwZmNkNzM2ZmY3ZGMzNTY1NmJhNDcw; kura_cloud_uid=15d500c8c525fc2821031462ca5c7cd0; odin_tt=aacad6dedcfc173c4b9642199bef2ce5e6b106a14f058ef08b6c1027d92f529897a0658ee79df54f3638fbb957fa4db71ba57772d7b0c32c5dd635fa92b613de; strategyABtestKey=%221737080660.037%22; _tea_utm_cache_4075=undefined; _tea_utm_cache_315365=undefined; FORCE_LOGIN=%7B%22videoConsumedRemainSeconds%22%3A180%2C%22isForcePopClose%22%3A1%7D; WallpaperGuide=%7B%22showTime%22%3A1737089490219%2C%22closeTime%22%3A0%2C%22showCount%22%3A3%2C%22cursor1%22%3A47%2C%22cursor2%22%3A14%2C%22hoverTime%22%3A1736409792142%7D; stream_player_status_params=%22%7B%5C%22is_auto_play%5C%22%3A0%2C%5C%22is_full_screen%5C%22%3A0%2C%5C%22is_full_webscreen%5C%22%3A0%2C%5C%22is_mute%5C%22%3A1%2C%5C%22is_speed%5C%22%3A1%2C%5C%22is_visible%5C%22%3A0%7D%22; stream_recommend_feed_params=%22%7B%5C%22cookie_enabled%5C%22%3Atrue%2C%5C%22screen_width%5C%22%3A1920%2C%5C%22screen_height%5C%22%3A1080%2C%5C%22browser_online%5C%22%3Atrue%2C%5C%22cpu_core_num%5C%22%3A12%2C%5C%22device_memory%5C%22%3A8%2C%5C%22downlink%5C%22%3A10%2C%5C%22effective_type%5C%22%3A%5C%224g%5C%22%2C%5C%22round_trip_time%5C%22%3A50%7D%22; bd_ticket_guard_client_data=eyJiZC10aWNrZXQtZ3VhcmQtdmVyc2lvbiI6MiwiYmQtdGlja2V0LWd1YXJkLWl0ZXJhdGlvbi12ZXJzaW9uIjoxLCJiZC10aWNrZXQtZ3VhcmQtcmVlLXB1YmxpYy1rZXkiOiJCQzRwQi9GVCtGZkpYTU01YjJBWXJUUGJIdmtJYmM5OVBYZmVXYlVKWm1ySHFRYlBCcW5JVXd4OUMyV3ZSUkJzVGNHQy94VnNZVTdsWTNPd1c5Qk04NE09IiwiYmQtdGlja2V0LWd1YXJkLXdlYi12ZXJzaW9uIjoyfQ%3D%3D; biz_trace_id=695d28b8; home_can_add_dy_2_desktop=%221%22; gfkadpd=315365,22350; IsDouyinActive=false; xgplayer_device_id=61595347671; xgplayer_user_id=695180662444; __live_version__=%221.1.2.7740%22; live_can_add_dy_2_desktop=%221%22; ttwid=1%7CjMZvmDSly5h6ELwmE9FI5212QkGoBXDWTPMZZPJOuEY%7C1737444787%7Cfc39140ae35fdc83527eebe52f8d228d4bda40af9019ac56ade26d6e88d2b6de; eos_s_token=CkLyeo38ina27TW0end6ORfbbeENp1oF3cXsSDla2itaQS89qlqqOPw42zN9p4YXyGJldhONLprFJLu06L5JL25dHPwaSQo8eiFMdGHg3xW4z4ZT2hduUNBC8pTku4cm2XPWpK9eB7lfwSWLJVsrF3J47WurtVmAtyJE2MP1HJJy6CStEOO15w0YweCp5w8iAQOD1/OJ"
	headers := kss.S{
		"User-Agent": reqs.GenRandomUa(),
		"Cookie":     cookie,
	}
	resp, err := client.Post(api, headers, nil, payload)
	if err != nil {
		log.Println(err)
	}
	jsonData, err := resp.JSON()
	if err != nil {
		log.Println(err)
	}
	data := kss.Nget(jsonData, []string{"data"}, nil)
	for _, one := range data.([]any) {
		if v, ok := one.(map[string]any); ok {
			fmt.Println(v["room_id"])
		}
	}
}
