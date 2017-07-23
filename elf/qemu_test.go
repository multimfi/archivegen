package elf

var qemufiles = map[string]ef{
	"/usr/bin/qemu-system-x86_64": {[]string{
		"libvirglrenderer.so.0",
		"libepoxy.so.0",
		"libdrm.so.2",
		"libgbm.so.1",
		"libX11.so.6",
		"libz.so.1",
		"libaio.so.1",
		"libpixman-1.so.0",
		"libutil.so.1",
		"libnuma.so.1",
		"libbluetooth.so.3",
		"libncursesw.so.6",
		"libbrlapi.so.0.6",
		"libasound.so.2",
		"libpulse.so.0",
		"libvdeplug.so.3",
		"libpng16.so.16",
		"libjpeg.so.8",
		"libsasl2.so.3",
		"libSDL2-2.0.so.0",
		"libvte-2.91.so.0",
		"libgtk-3.so.0",
		"libgdk-3.so.0",
		"libcairo.so.2",
		"libgdk_pixbuf-2.0.so.0",
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libgnutls.so.30",
		"libnettle.so.6",
		"liblzo2.so.2",
		"libsnappy.so.1",
		"libseccomp.so.2",
		"libspice-server.so.1",
		"libcacard.so.0",
		"libusb-1.0.so.0",
		"libusbredirparser.so.1",
		"libjemalloc.so.2",
		"libgmodule-2.0.so.0",
		"librt.so.1",
		"libm.so.6",
		"libgcc_s.so.1",
		"libpthread.so.0",
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/usr/lib/libvirglrenderer.so.0": {[]string{
		"libm.so.6",
		"libgbm.so.1",
		"libepoxy.so.0",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libm.so.6": {[]string{
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/usr/lib/libc.so.6": {[]string{
		"ld-linux-x86-64.so.2"}},

	"/usr/lib/ld-linux-x86-64.so.2": {},

	"/usr/lib/libgbm.so.1": {[]string{
		"libexpat.so.1",
		"libm.so.6",
		"libdl.so.2",
		"libwayland-client.so.0",
		"libwayland-server.so.0",
		"libdrm.so.2",
		"libc.so.6"}},

	"/usr/lib/libexpat.so.1": {[]string{
		"libc.so.6"}},

	"/usr/lib/libdl.so.2": {[]string{
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/usr/lib/libwayland-client.so.0": {[]string{
		"libffi.so.6",
		"librt.so.1",
		"libm.so.6",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libffi.so.6": {[]string{
		"libc.so.6"}},

	"/usr/lib/librt.so.1": {[]string{
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libpthread.so.0": {[]string{
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/usr/lib/libwayland-server.so.0": {[]string{
		"libffi.so.6",
		"librt.so.1",
		"libm.so.6",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libdrm.so.2": {[]string{
		"libm.so.6",
		"libc.so.6"}},

	"/usr/lib/libepoxy.so.0": {[]string{
		"libdl.so.2",
		"libc.so.6"}},

	"/usr/lib/libX11.so.6": {[]string{
		"libxcb.so.1",
		"libdl.so.2",
		"libc.so.6"}},

	"/usr/lib/libxcb.so.1": {[]string{
		"libXau.so.6",
		"libXdmcp.so.6",
		"libc.so.6"}},

	"/usr/lib/libXau.so.6": {[]string{
		"libc.so.6"}},

	"/usr/lib/libXdmcp.so.6": {[]string{
		"libc.so.6"}},

	"/usr/lib/libz.so.1": {[]string{
		"libc.so.6"}},

	"/usr/lib/libaio.so.1": {},

	"/usr/lib/libpixman-1.so.0": {[]string{
		"libm.so.6",
		"libpthread.so.0",
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/usr/lib/libutil.so.1": {[]string{
		"libc.so.6"}},

	"/usr/lib/libnuma.so.1": {[]string{
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/usr/lib/libbluetooth.so.3": {[]string{
		"libc.so.6"}},

	"/usr/lib/libncursesw.so.6": {[]string{
		"libc.so.6"}},

	"/usr/lib/libbrlapi.so.0.6": {[]string{
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libasound.so.2": {[]string{
		"libm.so.6",
		"libdl.so.2",
		"libpthread.so.0",
		"librt.so.1",
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/usr/lib/libpulse.so.0": {[]string{
		"libpulsecommon-10.0.so",
		"libdbus-1.so.3",
		"libpthread.so.0",
		"libdl.so.2",
		"libm.so.6",
		"libc.so.6"},
		{},
		[]string{"/usr/lib/pulseaudio"}},

	"/usr/lib/pulseaudio/libpulsecommon-10.0.so": {[]string{
		"libxcb.so.1",
		"libsystemd.so.0",
		"libsndfile.so.1",
		"libasyncns.so.0",
		"libdbus-1.so.3",
		"libpthread.so.0",
		"librt.so.1",
		"libdl.so.2",
		"libm.so.6",
		"libc.so.6"}},

	"/usr/lib/libsystemd.so.0": {[]string{
		"libresolv.so.2",
		"libcap.so.2",
		"librt.so.1",
		"liblzma.so.5",
		"liblz4.so.1",
		"libgcrypt.so.20",
		"libgpg-error.so.0",
		"libpthread.so.0",
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/usr/lib/libresolv.so.2": {[]string{
		"libc.so.6"}},

	"/usr/lib/libcap.so.2": {[]string{
		"libc.so.6"}},

	"/usr/lib/liblzma.so.5": {[]string{
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/liblz4.so.1": {[]string{
		"libc.so.6"}},

	"/usr/lib/libgcrypt.so.20": {[]string{
		"libgpg-error.so.0",
		"libc.so.6"}},

	"/usr/lib/libgpg-error.so.0": {[]string{
		"libc.so.6"}},

	"/usr/lib/libsndfile.so.1": {[]string{
		"libFLAC.so.8",
		"libogg.so.0",
		"libvorbis.so.0",
		"libvorbisenc.so.2",
		"libm.so.6",
		"libc.so.6"}},

	"/usr/lib/libFLAC.so.8": {[]string{
		"libogg.so.0",
		"libm.so.6",
		"libc.so.6"}},

	"/usr/lib/libogg.so.0": {[]string{
		"libc.so.6"}},

	"/usr/lib/libvorbis.so.0": {[]string{
		"libm.so.6",
		"libogg.so.0",
		"libc.so.6"}},

	"/usr/lib/libvorbisenc.so.2": {[]string{
		"libvorbis.so.0",
		"libm.so.6",
		"libogg.so.0",
		"libc.so.6"}},

	"/usr/lib/libasyncns.so.0": {[]string{
		"libresolv.so.2",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libdbus-1.so.3": {[]string{
		"libpthread.so.0",
		"libsystemd.so.0",
		"libc.so.6"}},

	"/usr/lib/libvdeplug.so.3": {[]string{
		"libdl.so.2",
		"libc.so.6"}},

	"/usr/lib/libpng16.so.16": {[]string{
		"libz.so.1",
		"libm.so.6",
		"libc.so.6"}},

	"/usr/lib/libjpeg.so.8": {[]string{
		"libc.so.6"}},

	"/usr/lib/libsasl2.so.3": {[]string{
		"libdl.so.2",
		"libresolv.so.2",
		"libc.so.6"}},

	"/usr/lib/libSDL2-2.0.so.0": {[]string{
		"libm.so.6",
		"libdl.so.2",
		"libpthread.so.0",
		"librt.so.1",
		"libc.so.6"}},

	"/usr/lib/libvte-2.91.so.0": {[]string{
		"libgtk-3.so.0",
		"libgdk-3.so.0",
		"libpangocairo-1.0.so.0",
		"libpango-1.0.so.0",
		"libatk-1.0.so.0",
		"libcairo-gobject.so.2",
		"libcairo.so.2",
		"libgdk_pixbuf-2.0.so.0",
		"libgio-2.0.so.0",
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libz.so.1",
		"libpcre2-8.so.0",
		"libgnutls.so.30",
		"libstdc++.so.6",
		"libm.so.6",
		"libc.so.6",
		"libgcc_s.so.1"}},

	"/usr/lib/libgtk-3.so.0": {[]string{
		"libgdk-3.so.0",
		"libgmodule-2.0.so.0",
		"libpangocairo-1.0.so.0",
		"libX11.so.6",
		"libXi.so.6",
		"libXfixes.so.3",
		"libcairo-gobject.so.2",
		"libcairo.so.2",
		"libgdk_pixbuf-2.0.so.0",
		"libatk-1.0.so.0",
		"libatk-bridge-2.0.so.0",
		"libepoxy.so.0",
		"libpangoft2-1.0.so.0",
		"libpango-1.0.so.0",
		"libfontconfig.so.1",
		"libgio-2.0.so.0",
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libm.so.6",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libgdk-3.so.0": {[]string{
		"libpangocairo-1.0.so.0",
		"libpango-1.0.so.0",
		"libgdk_pixbuf-2.0.so.0",
		"libcairo-gobject.so.2",
		"libgio-2.0.so.0",
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libfontconfig.so.1",
		"libXinerama.so.1",
		"libXi.so.6",
		"libXrandr.so.2",
		"libXcursor.so.1",
		"libXcomposite.so.1",
		"libXdamage.so.1",
		"libXfixes.so.3",
		"libxkbcommon.so.0",
		"libwayland-cursor.so.0",
		"libwayland-egl.so.1",
		"libwayland-client.so.0",
		"libX11.so.6",
		"libXext.so.6",
		"libcairo.so.2",
		"libepoxy.so.0",
		"libm.so.6",
		"librt.so.1",
		"libc.so.6"}},

	"/usr/lib/libpangocairo-1.0.so.0": {[]string{
		"libpango-1.0.so.0",
		"libpangoft2-1.0.so.0",
		"libm.so.6",
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libfontconfig.so.1",
		"libfreetype.so.6",
		"libcairo.so.2",
		"libc.so.6"}},

	"/usr/lib/libpango-1.0.so.0": {[]string{
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libm.so.6",
		"libthai.so.0",
		"libc.so.6"}},

	"/usr/lib/libgobject-2.0.so.0": {[]string{
		"libglib-2.0.so.0",
		"libffi.so.6",
		"libc.so.6"}},

	"/usr/lib/libglib-2.0.so.0": {[]string{
		"libpcre.so.1",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libpcre.so.1": {[]string{
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libthai.so.0": {[]string{
		"libdatrie.so.1",
		"libc.so.6"}},

	"/usr/lib/libdatrie.so.1": {[]string{
		"libc.so.6"}},

	"/usr/lib/libpangoft2-1.0.so.0": {[]string{
		"libpango-1.0.so.0",
		"libm.so.6",
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libharfbuzz.so.0",
		"libfontconfig.so.1",
		"libfreetype.so.6",
		"libc.so.6"}},

	"/usr/lib/libharfbuzz.so.0": {[]string{
		"libglib-2.0.so.0",
		"libfreetype.so.6",
		"libgraphite2.so.3",
		"libc.so.6"}},

	"/usr/lib/libfreetype.so.6": {[]string{
		"libz.so.1",
		"libbz2.so.1.0",
		"libpng16.so.16",
		"libharfbuzz.so.0",
		"libc.so.6"}},

	"/usr/lib/libbz2.so.1.0": {[]string{
		"libc.so.6"}},

	"/usr/lib/libgraphite2.so.3": {[]string{
		"libc.so.6"}},

	"/usr/lib/libfontconfig.so.1": {[]string{
		"libfreetype.so.6",
		"libexpat.so.1",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libcairo.so.2": {[]string{
		"libpthread.so.0",
		"libpixman-1.so.0",
		"libfontconfig.so.1",
		"libfreetype.so.6",
		"libEGL.so.1",
		"libdl.so.2",
		"libpng16.so.16",
		"libxcb-shm.so.0",
		"libxcb.so.1",
		"libxcb-render.so.0",
		"libXrender.so.1",
		"libX11.so.6",
		"libXext.so.6",
		"libz.so.1",
		"libGL.so.1",
		"librt.so.1",
		"libm.so.6",
		"libc.so.6"}},

	"/usr/lib/libEGL.so.1": {[]string{
		"libm.so.6",
		"libGLdispatch.so.0",
		"libdl.so.2",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libGLdispatch.so.0": {[]string{
		"libdl.so.2",
		"libc.so.6"}},

	"/usr/lib/libxcb-shm.so.0": {[]string{
		"libxcb.so.1",
		"libXau.so.6",
		"libXdmcp.so.6",
		"libc.so.6"}},

	"/usr/lib/libxcb-render.so.0": {[]string{
		"libxcb.so.1",
		"libXau.so.6",
		"libXdmcp.so.6",
		"libc.so.6"}},

	"/usr/lib/libXrender.so.1": {[]string{
		"libX11.so.6",
		"libc.so.6"}},

	"/usr/lib/libXext.so.6": {[]string{
		"libX11.so.6",
		"libc.so.6"}},

	"/usr/lib/libGL.so.1": {[]string{
		"libGLX.so.0",
		"libX11.so.6",
		"libXext.so.6",
		"libGLdispatch.so.0",
		"libdl.so.2",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libGLX.so.0": {[]string{
		"libX11.so.6",
		"libXext.so.6",
		"libGLdispatch.so.0",
		"libdl.so.2",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libgdk_pixbuf-2.0.so.0": {[]string{
		"libgmodule-2.0.so.0",
		"libgio-2.0.so.0",
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libpng16.so.16",
		"libm.so.6",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libgmodule-2.0.so.0": {[]string{
		"libdl.so.2",
		"libglib-2.0.so.0",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libgio-2.0.so.0": {[]string{
		"libgobject-2.0.so.0",
		"libgmodule-2.0.so.0",
		"libglib-2.0.so.0",
		"libpthread.so.0",
		"libz.so.1",
		"libresolv.so.2",
		"libmount.so.1",
		"libc.so.6"}},

	"/usr/lib/libmount.so.1": {[]string{
		"libblkid.so.1",
		"libuuid.so.1",
		"librt.so.1",
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/usr/lib/libblkid.so.1": {[]string{
		"libuuid.so.1",
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/usr/lib/libuuid.so.1": {[]string{
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/usr/lib/libcairo-gobject.so.2": {[]string{
		"libcairo.so.2",
		"libpthread.so.0",
		"libpixman-1.so.0",
		"libfontconfig.so.1",
		"libfreetype.so.6",
		"libEGL.so.1",
		"libdl.so.2",
		"libpng16.so.16",
		"libxcb-shm.so.0",
		"libxcb.so.1",
		"libxcb-render.so.0",
		"libXrender.so.1",
		"libX11.so.6",
		"libXext.so.6",
		"libz.so.1",
		"libGL.so.1",
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"librt.so.1",
		"libm.so.6",
		"libc.so.6"}},

	"/usr/lib/libXinerama.so.1": {[]string{
		"libX11.so.6",
		"libXext.so.6",
		"libc.so.6"}},

	"/usr/lib/libXi.so.6": {[]string{
		"libX11.so.6",
		"libXext.so.6",
		"libc.so.6"}},

	"/usr/lib/libXrandr.so.2": {[]string{
		"libXext.so.6",
		"libXrender.so.1",
		"libX11.so.6",
		"libc.so.6"}},

	"/usr/lib/libXcursor.so.1": {[]string{
		"libXrender.so.1",
		"libXfixes.so.3",
		"libX11.so.6",
		"libc.so.6"}},

	"/usr/lib/libXfixes.so.3": {[]string{
		"libX11.so.6",
		"libc.so.6"}},

	"/usr/lib/libXcomposite.so.1": {[]string{
		"libX11.so.6",
		"libc.so.6"}},

	"/usr/lib/libXdamage.so.1": {[]string{
		"libXfixes.so.3",
		"libX11.so.6",
		"libc.so.6"}},

	"/usr/lib/libxkbcommon.so.0": {[]string{
		"libc.so.6"}},

	"/usr/lib/libwayland-cursor.so.0": {[]string{
		"libwayland-client.so.0",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libwayland-egl.so.1": {[]string{
		"libc.so.6"}},

	"/usr/lib/libatk-1.0.so.0": {[]string{
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libc.so.6"}},

	"/usr/lib/libatk-bridge-2.0.so.0": {[]string{
		"libatk-1.0.so.0",
		"libgobject-2.0.so.0",
		"libatspi.so.0",
		"libdbus-1.so.3",
		"libglib-2.0.so.0",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libatspi.so.0": {[]string{
		"libgobject-2.0.so.0",
		"libX11.so.6",
		"libdbus-1.so.3",
		"libglib-2.0.so.0",
		"libc.so.6"}},

	"/usr/lib/libpcre2-8.so.0": {[]string{
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libgnutls.so.30": {[]string{
		"libz.so.1",
		"libp11-kit.so.0",
		"libunistring.so.2",
		"libtasn1.so.6",
		"libnettle.so.6",
		"libhogweed.so.4",
		"libgmp.so.10",
		"libc.so.6"}},

	"/usr/lib/libp11-kit.so.0": {[]string{
		"libffi.so.6",
		"libdl.so.2",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libunistring.so.2": {[]string{
		"libc.so.6"}},

	"/usr/lib/libtasn1.so.6": {[]string{
		"libc.so.6"}},

	"/usr/lib/libnettle.so.6": {[]string{
		"libc.so.6"}},

	"/usr/lib/libhogweed.so.4": {[]string{
		"libnettle.so.6",
		"libgmp.so.10",
		"libc.so.6"}},

	"/usr/lib/libgmp.so.10": {[]string{
		"libc.so.6"}},

	"/usr/lib/libstdc++.so.6": {[]string{
		"libm.so.6",
		"libc.so.6",
		"ld-linux-x86-64.so.2",
		"libgcc_s.so.1"}},

	"/usr/lib/libgcc_s.so.1": {[]string{
		"libc.so.6"}},

	"/usr/lib/liblzo2.so.2": {[]string{
		"libc.so.6"}},

	"/usr/lib/libsnappy.so.1": {[]string{
		"libstdc++.so.6",
		"libm.so.6",
		"libc.so.6",
		"libgcc_s.so.1"}},

	"/usr/lib/libseccomp.so.2": {[]string{
		"libc.so.6"}},

	"/usr/lib/libspice-server.so.1": {[]string{
		"libcelt051.so.0",
		"libgio-2.0.so.0",
		"libgobject-2.0.so.0",
		"libglib-2.0.so.0",
		"libjpeg.so.8",
		"libpixman-1.so.0",
		"libsasl2.so.3",
		"libssl.so.1.1",
		"libcrypto.so.1.1",
		"libz.so.1",
		"libm.so.6",
		"librt.so.1",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libcelt051.so.0": {[]string{
		"libm.so.6",
		"libc.so.6"}},

	"/usr/lib/libssl.so.1.1": {[]string{
		"libcrypto.so.1.1",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libcrypto.so.1.1": {[]string{
		"libdl.so.2",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libcacard.so.0": {[]string{
		"libglib-2.0.so.0",
		"libssl3.so",
		"libsmime3.so",
		"libnss3.so",
		"libnssutil3.so",
		"libplds4.so",
		"libplc4.so",
		"libnspr4.so",
		"libc.so.6"}},

	"/usr/lib/libssl3.so": {[]string{
		"libnss3.so",
		"libnssutil3.so",
		"libpthread.so.0",
		"libc.so.6",
		"libz.so.1",
		"libplc4.so",
		"libnspr4.so"}},

	"/usr/lib/libnss3.so": {[]string{
		"libnssutil3.so",
		"libc.so.6",
		"libplds4.so",
		"libplc4.so",
		"libnspr4.so"}},

	"/usr/lib/libnssutil3.so": {[]string{
		"libpthread.so.0",
		"libc.so.6",
		"libplds4.so",
		"libplc4.so",
		"libnspr4.so"}},

	"/usr/lib/libplds4.so": {[]string{
		"libnspr4.so",
		"libc.so.6"}},

	"/usr/lib/libnspr4.so": {[]string{
		"libpthread.so.0",
		"libdl.so.2",
		"librt.so.1",
		"libc.so.6"}},

	"/usr/lib/libplc4.so": {[]string{
		"libnspr4.so",
		"libc.so.6"}},

	"/usr/lib/libsmime3.so": {[]string{
		"libnss3.so",
		"libnssutil3.so",
		"libc.so.6",
		"libplds4.so",
		"libplc4.so",
		"libnspr4.so"}},

	"/usr/lib/libusb-1.0.so.0": {[]string{
		"libudev.so.1",
		"libpthread.so.0",
		"libc.so.6"}},

	"/usr/lib/libudev.so.1": {[]string{
		"libresolv.so.2",
		"libcap.so.2",
		"librt.so.1",
		"libpthread.so.0",
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},

	"/usr/lib/libusbredirparser.so.1": {[]string{
		"libc.so.6"}},

	"/usr/lib/libjemalloc.so.2": {[]string{
		"libstdc++.so.6",
		"libpthread.so.0",
		"libdl.so.2",
		"libgcc_s.so.1",
		"libc.so.6",
		"ld-linux-x86-64.so.2"}},
}

var qemuresolved = map[string]string{
	"libgpg-error.so.0":      "/usr/lib/libgpg-error.so.0",
	"libxcb-shm.so.0":        "/usr/lib/libxcb-shm.so.0",
	"libcrypto.so.1.1":       "/usr/lib/libcrypto.so.1.1",
	"libvdeplug.so.3":        "/usr/lib/libvdeplug.so.3",
	"libEGL.so.1":            "/usr/lib/libEGL.so.1",
	"libunistring.so.2":      "/usr/lib/libunistring.so.2",
	"libnettle.so.6":         "/usr/lib/libnettle.so.6",
	"ld-linux-x86-64.so.2":   "/usr/lib/ld-linux-x86-64.so.2",
	"libpng16.so.16":         "/usr/lib/libpng16.so.16",
	"libglib-2.0.so.0":       "/usr/lib/libglib-2.0.so.0",
	"libgmodule-2.0.so.0":    "/usr/lib/libgmodule-2.0.so.0",
	"libXi.so.6":             "/usr/lib/libXi.so.6",
	"libmount.so.1":          "/usr/lib/libmount.so.1",
	"libnss3.so":             "/usr/lib/libnss3.so",
	"libdl.so.2":             "/usr/lib/libdl.so.2",
	"libdrm.so.2":            "/usr/lib/libdrm.so.2",
	"libFLAC.so.8":           "/usr/lib/libFLAC.so.8",
	"libthai.so.0":           "/usr/lib/libthai.so.0",
	"libbz2.so.1.0":          "/usr/lib/libbz2.so.1.0",
	"libcairo-gobject.so.2":  "/usr/lib/libcairo-gobject.so.2",
	"libwayland-server.so.0": "/usr/lib/libwayland-server.so.0",
	"libXext.so.6":           "/usr/lib/libXext.so.6",
	"libblkid.so.1":          "/usr/lib/libblkid.so.1",
	"libatspi.so.0":          "/usr/lib/libatspi.so.0",
	"libasound.so.2":         "/usr/lib/libasound.so.2",
	"libpango-1.0.so.0":      "/usr/lib/libpango-1.0.so.0",
	"libXrender.so.1":        "/usr/lib/libXrender.so.1",
	"libgio-2.0.so.0":        "/usr/lib/libgio-2.0.so.0",
	"libatk-1.0.so.0":        "/usr/lib/libatk-1.0.so.0",
	"libsnappy.so.1":         "/usr/lib/libsnappy.so.1",
	"libssl.so.1.1":          "/usr/lib/libssl.so.1.1",
	"libXau.so.6":            "/usr/lib/libXau.so.6",
	"libfontconfig.so.1":     "/usr/lib/libfontconfig.so.1",
	"libXcomposite.so.1":     "/usr/lib/libXcomposite.so.1",
	"libXdamage.so.1":        "/usr/lib/libXdamage.so.1",
	"libwayland-cursor.so.0": "/usr/lib/libwayland-cursor.so.0",
	"libGL.so.1":             "/usr/lib/libGL.so.1",
	"libuuid.so.1":           "/usr/lib/libuuid.so.1",
	"libwayland-egl.so.1":    "/usr/lib/libwayland-egl.so.1",
	"libsndfile.so.1":        "/usr/lib/libsndfile.so.1",
	"libvorbis.so.0":         "/usr/lib/libvorbis.so.0",
	"libdatrie.so.1":         "/usr/lib/libdatrie.so.1",
	"libGLdispatch.so.0":     "/usr/lib/libGLdispatch.so.0",
	"libxcb-render.so.0":     "/usr/lib/libxcb-render.so.0",
	"libplc4.so":             "/usr/lib/libplc4.so",
	"libharfbuzz.so.0":       "/usr/lib/libharfbuzz.so.0",
	"libplds4.so":            "/usr/lib/libplds4.so",
	"libpixman-1.so.0":       "/usr/lib/libpixman-1.so.0",
	"libjpeg.so.8":           "/usr/lib/libjpeg.so.8",
	"libgcc_s.so.1":          "/usr/lib/libgcc_s.so.1",
	"liblz4.so.1":            "/usr/lib/liblz4.so.1",
	"libogg.so.0":            "/usr/lib/libogg.so.0",
	"libGLX.so.0":            "/usr/lib/libGLX.so.0",
	"libXrandr.so.2":         "/usr/lib/libXrandr.so.2",
	"libXcursor.so.1":        "/usr/lib/libXcursor.so.1",
	"libz.so.1":              "/usr/lib/libz.so.1",
	"libbluetooth.so.3":      "/usr/lib/libbluetooth.so.3",
	"liblzma.so.5":           "/usr/lib/liblzma.so.5",
	"libsasl2.so.3":          "/usr/lib/libsasl2.so.3",
	"libpangocairo-1.0.so.0": "/usr/lib/libpangocairo-1.0.so.0",
	"libudev.so.1":           "/usr/lib/libudev.so.1",
	"libm.so.6":              "/usr/lib/libm.so.6",
	"libaio.so.1":            "/usr/lib/libaio.so.1",
	"libfreetype.so.6":       "/usr/lib/libfreetype.so.6",
	"libatk-bridge-2.0.so.0": "/usr/lib/libatk-bridge-2.0.so.0",
	"libtasn1.so.6":          "/usr/lib/libtasn1.so.6",
	"liblzo2.so.2":           "/usr/lib/liblzo2.so.2",
	"libXdmcp.so.6":          "/usr/lib/libXdmcp.so.6",
	"libgdk_pixbuf-2.0.so.0": "/usr/lib/libgdk_pixbuf-2.0.so.0",
	"libpulsecommon-10.0.so": "/usr/lib/pulseaudio/libpulsecommon-10.0.so",
	"libsystemd.so.0":        "/usr/lib/libsystemd.so.0",
	"libhogweed.so.4":        "/usr/lib/libhogweed.so.4",
	"libssl3.so":             "/usr/lib/libssl3.so",
	"libgraphite2.so.3":      "/usr/lib/libgraphite2.so.3",
	"libpcre2-8.so.0":        "/usr/lib/libpcre2-8.so.0",
	"libgmp.so.10":           "/usr/lib/libgmp.so.10",
	"libwayland-client.so.0": "/usr/lib/libwayland-client.so.0",
	"libffi.so.6":            "/usr/lib/libffi.so.6",
	"librt.so.1":             "/usr/lib/librt.so.1",
	"libutil.so.1":           "/usr/lib/libutil.so.1",
	"libnuma.so.1":           "/usr/lib/libnuma.so.1",
	"libjemalloc.so.2":       "/usr/lib/libjemalloc.so.2",
	"libcairo.so.2":          "/usr/lib/libcairo.so.2",
	"libstdc++.so.6":         "/usr/lib/libstdc++.so.6",
	"libusbredirparser.so.1": "/usr/lib/libusbredirparser.so.1",
	"libc.so.6":              "/usr/lib/libc.so.6",
	"libncursesw.so.6":       "/usr/lib/libncursesw.so.6",
	"libbrlapi.so.0.6":       "/usr/lib/libbrlapi.so.0.6",
	"libasyncns.so.0":        "/usr/lib/libasyncns.so.0",
	"libgobject-2.0.so.0":    "/usr/lib/libgobject-2.0.so.0",
	"libgcrypt.so.20":        "/usr/lib/libgcrypt.so.20",
	"libp11-kit.so.0":        "/usr/lib/libp11-kit.so.0",
	"libspice-server.so.1":   "/usr/lib/libspice-server.so.1",
	"libpulse.so.0":          "/usr/lib/libpulse.so.0",
	"libgnutls.so.30":        "/usr/lib/libgnutls.so.30",
	"libseccomp.so.2":        "/usr/lib/libseccomp.so.2",
	"libXfixes.so.3":         "/usr/lib/libXfixes.so.3",
	"qemu-system-x86_64":     "/usr/bin/qemu-system-x86_64",
	"libvirglrenderer.so.0":  "/usr/lib/libvirglrenderer.so.0",
	"libexpat.so.1":          "/usr/lib/libexpat.so.1",
	"libpthread.so.0":        "/usr/lib/libpthread.so.0",
	"libresolv.so.2":         "/usr/lib/libresolv.so.2",
	"libcelt051.so.0":        "/usr/lib/libcelt051.so.0",
	"libxcb.so.1":            "/usr/lib/libxcb.so.1",
	"libvorbisenc.so.2":      "/usr/lib/libvorbisenc.so.2",
	"libgdk-3.so.0":          "/usr/lib/libgdk-3.so.0",
	"libpcre.so.1":           "/usr/lib/libpcre.so.1",
	"libXinerama.so.1":       "/usr/lib/libXinerama.so.1",
	"libcap.so.2":            "/usr/lib/libcap.so.2",
	"libdbus-1.so.3":         "/usr/lib/libdbus-1.so.3",
	"libSDL2-2.0.so.0":       "/usr/lib/libSDL2-2.0.so.0",
	"libcacard.so.0":         "/usr/lib/libcacard.so.0",
	"libnssutil3.so":         "/usr/lib/libnssutil3.so",
	"libgbm.so.1":            "/usr/lib/libgbm.so.1",
	"libsmime3.so":           "/usr/lib/libsmime3.so",
	"libxkbcommon.so.0":      "/usr/lib/libxkbcommon.so.0",
	"libusb-1.0.so.0":        "/usr/lib/libusb-1.0.so.0",
	"libgtk-3.so.0":          "/usr/lib/libgtk-3.so.0",
	"libepoxy.so.0":          "/usr/lib/libepoxy.so.0",
	"libX11.so.6":            "/usr/lib/libX11.so.6",
	"libvte-2.91.so.0":       "/usr/lib/libvte-2.91.so.0",
	"libpangoft2-1.0.so.0":   "/usr/lib/libpangoft2-1.0.so.0",
	"libnspr4.so":            "/usr/lib/libnspr4.so",
}
