package mumbleclient

type MumbleVersion uint32

func EncodeMumbleVersion(major uint16, minor uint8, patch uint8) MumbleVersion {
	ver := uint32(0)
	ver |= uint32(major) << 16
	ver |= uint32(minor) << 8
	ver |= uint32(patch)

	return MumbleVersion(ver)
}
