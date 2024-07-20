from pydub import AudioSegment

def pcm_to_mp3(pcm_file_path, mp3_file_path, sample_rate=16000, channels=1):
    # Load PCM file
    audio = AudioSegment.from_raw(
        pcm_file_path,
        sample_width=2,  # 16-bit PCM
        frame_rate=sample_rate,
        channels=channels
    )

    # Export as MP3
    audio.export(mp3_file_path, format="mp3")

# 示例用法
pcm_file_path = "output.pcm"
mp3_file_path = "output.mp3"
pcm_to_mp3(pcm_file_path, mp3_file_path)
