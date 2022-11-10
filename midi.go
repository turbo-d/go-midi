package midi

// Channel voice message status
const (
	noteOffStatus         byte = 0x80
	noteOnStatus          byte = 0x90
	polyKeyPressureStatus byte = 0xA0
	controlChangeStatus   byte = 0xB0
	programChangeStatus   byte = 0xC0
	channelPressureStatus byte = 0xD0
	pitchBendChangeStatus byte = 0xE0
)

// Channel mode message status
const (
	selectChannelModeStatus byte = 0xB0
)

// System common message status
const (
	midiTimeCodeQtrFrameStatus byte = 0xF1
	songPositionPtrStatus      byte = 0xF2
	songSelectStatus           byte = 0xF3
	tuneRequestStatus          byte = 0xF6
)

// System real time message status
const (
	timingClockStatus   byte = 0xF8
	startStatus         byte = 0xFA
	continueStatus      byte = 0xFB
	stopStatus          byte = 0xFC
	activeSensingStatus byte = 0xFE
	systemResetStatus   byte = 0xFF
)

// System exclusive message status
const (
	soxStatus byte = 0xF0
	eoxStatus byte = 0xF7
)

// Control numbers
const (
	// 0x00 - 0x1F: MSB of 14-bit controls
	bankSelectControl                byte = 0x00
	modWheelControl                  byte = 0x01
	breathControllerControl          byte = 0x02
	footControllerControl            byte = 0x04
	portamentoTimeControl            byte = 0x05
	dataEntryMSB                     byte = 0x06
	channelVolumeControl             byte = 0x07
	balanceControl                   byte = 0x08
	panControl                       byte = 0x0A
	expressionControllerControl      byte = 0x0B
	effectControl1Control            byte = 0x0C
	effectControl2Control            byte = 0x0D
	generalPurposeController1Control byte = 0x10
	generalPurposeController2Control byte = 0x11
	generalPurposeController3Control byte = 0x12
	generalPurposeController4Control byte = 0x13
	// 0x20 - 0x3F: Optional LSB of 14-bit controls
	lsb0Control  byte = 0x20
	lsb1Control  byte = 0x21
	lsb2Control  byte = 0x22
	lsb3Control  byte = 0x23
	lsb4Control  byte = 0x24
	lsb5Control  byte = 0x25
	lsb6Control  byte = 0x26
	lsb7Control  byte = 0x27
	lsb8Control  byte = 0x28
	lsb9Control  byte = 0x29
	lsb10Control byte = 0x2A
	lsb11Control byte = 0x2B
	lsb12Control byte = 0x2C
	lsb13Control byte = 0x2D
	lsb14Control byte = 0x2E
	lsb15Control byte = 0x2F
	lsb16Control byte = 0x30
	lsb17Control byte = 0x31
	lsb18Control byte = 0x32
	lsb19Control byte = 0x33
	lsb20Control byte = 0x34
	lsb21Control byte = 0x35
	lsb22Control byte = 0x36
	lsb23Control byte = 0x37
	lsb24Control byte = 0x38
	lsb25Control byte = 0x39
	lsb26Control byte = 0x3A
	lsb27Control byte = 0x3B
	lsb28Control byte = 0x3C
	lsb29Control byte = 0x3D
	lsb30Control byte = 0x3E
	lsb31Control byte = 0x3F
	// 0x40 - 0x45: 7-bit controls for switched funtions
	damperPedalControl      byte = 0x40
	portamentoOnOffControl  byte = 0x41
	sostenutoControl        byte = 0x42
	softPedalControl        byte = 0x43
	legatoFootswitchControl byte = 0x44
	hold2Control            byte = 0x45
	// 0x46 - 0x5F: 7-bit controls for effect depth
	soundController1Control          byte = 0x46
	soundController2Control          byte = 0x47
	soundController3Control          byte = 0x48
	soundController4Control          byte = 0x49
	soundController5Control          byte = 0x4A
	soundController6Control          byte = 0x4B
	soundController7Control          byte = 0x4C
	soundController8Control          byte = 0x4D
	soundController9Control          byte = 0x4E
	soundController10Control         byte = 0x4F
	generalPurposeController5Control byte = 0x50
	generalPurposeController6Control byte = 0x51
	generalPurposeController7Control byte = 0x52
	generalPurposeController8Control byte = 0x53
	portamentoControlControl         byte = 0x54
	effects1DepthControl             byte = 0x5B
	effects2DepthControl             byte = 0x5C
	effects3DepthControl             byte = 0x5D
	effects4DepthControl             byte = 0x5E
	effects5DepthControl             byte = 0x5F
	// 0x60 - 0x65: Inc/Dec and Parameter numbers
	dataIncrementControl                   byte = 0x60
	dataDecrementControl                   byte = 0x61
	nonRegisteredParameterNumberLSBControl byte = 0x62
	nonRegisteredParameterNumberMSBControl byte = 0x63
	registeredParameterNumberLSBControl    byte = 0x64
	registeredParameterNumberMSBControl    byte = 0x65
)

const (
	defaultVelocity byte = 0x40
	onValue         byte = 0x7F
	offValue        byte = 0x00
)

func isControlOn(ctrl byte) bool {
	if ctrl > 127 {
		// TODO: Error
	} else if ctrl < 64 {
		return false
	}
	return true
}

type midiEvent interface {
	isSysex() bool
}

type msg0 struct {
	status byte
}

func (m msg0) isSysex() bool {
	return false
}

type msg1 struct {
	status byte
	data0  byte
}

func (m msg1) isSysex() bool {
	return false
}

type msg2 struct {
	status byte
	data0  byte
	data1  byte
}

func (m msg2) isSysex() bool {
	return false
}

type Msg struct {
	status byte
	data   []byte
}

// TODO: Limit the allowable channels by creating a type, or do error checking in each factory method
// TODO: Implement Running Status mode optimizations on transmitting side
// TODO: Capability to do conversion to logarithmic function and back for velocity (or any relevant) inputs
// TODO: If MSB & LSB of CC are sent, then in subsequent fine tunings only the LSB must be sent. Who stores the state of the MSB, the app or the framework?

// Channel voice messages
func createNoteOff(ch, note, vel int) (*Msg, error) {
	// makeChannel(int) (byte, error)
	if ch < 1 || ch > 16 {
		// TODO: Error
	}
	// makeDataByte(int) (byte, error)
	if note < 0 || note > 127 {
		// TODO: Error
	}
	if vel < 0 || vel > 127 {
		// TODO: Error
	}

	return &Msg{
		status: noteOffStatus | byte(ch),
		data:   []byte{byte(note), byte(vel)},
	}, nil
}

func createNoteOffDefaultVelocity(ch byte, note byte) midiEvent {
	return &msg2{
		status: noteOffStatus | ch,
		data0:  note,
		data1:  defaultVelocity,
	}
}

func createNoteOffRunningStatus(ch byte, note byte) midiEvent {
	return &msg2{
		status: noteOnStatus | ch,
		data0:  note,
		data1:  0x00,
	}
}

func createNoteOn(ch byte, note byte, vel byte) midiEvent {
	return &msg2{
		status: noteOnStatus | ch,
		data0:  note,
		data1:  vel,
	}
}

func createNoteOnDefaultVelocity(ch byte, note byte) midiEvent {
	return &msg2{
		status: noteOnStatus | ch,
		data0:  note,
		data1:  defaultVelocity,
	}
}

func createPolyKeyPressure(ch byte, note byte, pressure byte) midiEvent {
	return &msg2{
		status: polyKeyPressureStatus | ch,
		data0:  note,
		data1:  pressure,
	}
}

func createControlChange(ch byte, ctrl byte, val byte) midiEvent {
	return &msg2{
		status: controlChangeStatus | ch,
		data0:  ctrl,
		data1:  val,
	}
}

func createProgramChange(ch byte, program byte) midiEvent {
	return &msg1{
		status: programChangeStatus | ch,
		data0:  program,
	}
}

func createChannelPressure(ch byte, pressure byte) midiEvent {
	return &msg1{
		status: channelPressureStatus | ch,
		data0:  pressure,
	}
}

func createPitchBendChange(ch byte, lsb byte, msb byte) midiEvent {
	return &msg2{
		status: pitchBendChangeStatus | ch,
		data0:  lsb,
		data1:  msb,
	}
}

// System common messages
func createMIDITimeCodeQtrFrame(msgType int, values int) midiEvent {
	return &msg1{
		status: midiTimeCodeQtrFrameStatus,
		data0:  0x00 | byte((msgType&0x7)<<4) | byte(values&0xF),
	}
}

func createSongPositionPointer(l byte, h byte) midiEvent {
	return &msg2{
		status: songPositionPtrStatus,
		data0:  0x00 | byte(l&0x7F),
		data1:  0x00 | byte(h&0x7F),
	}
}

func createSongSelect(song byte) midiEvent {
	return &msg1{
		status: songSelectStatus,
		data0:  0x00 | byte(song&0x7F),
	}
}

func createTuneRequest() midiEvent {
	return &msg0{
		status: tuneRequestStatus,
	}
}
