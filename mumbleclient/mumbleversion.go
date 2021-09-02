package mumbleclient

type MumbleVersion uint32

var CurrentMumbleVersion = EncodeMumbleVersion(1, 3, 4)

func EncodeMumbleVersion(major uint16, minor uint8, patch uint8) MumbleVersion {
	ver := uint32(0)
	ver |= uint32(major) << 16
	ver |= uint32(minor) << 8
	ver |= uint32(patch)

	return MumbleVersion(ver)
}
