# syz-crash-reporter

syz-crash-reporter is a simple plugin of syzkaller, aiming to add debug info into a simple crash log.

for example, take the input crash log:
```
ioctl$sock_SIOCDELRT(r0, 0x890c, &(0x7f0000001880)={0x0, @ax25={0x3, @default, 0x3}, @nl=@kern={0x10, 0x0, 0x0, 0x1000}, @sco={0x1f, @fixed={[], 0x11}}, 0x2006, 0x0, 0x0, 0x0, 0xe, 0x0, 0x80, 0x7, 0x6})
syz_open_dev$usbmon(&(0x7f0000001900)='/dev/usbmon#\x00', 0x8001, 0x8000)

[   57.570255][ T6164] ------------[ cut here ]------------
[   57.571363][ T6164] WARNING: CPU: 0 PID: 6164 at kernel/cgroup/cgroup.c:3111 cgroup_apply_control_disable+0x400/0x4d0
[   57.573581][ T6164] Kernel panic - not syncing: panic_on_warn set ...
[   57.574781][ T6164] CPU: 0 PID: 6164 Comm: syz-executor.1 Not tainted 5.7.0 #1
[   57.576393][ T6164] Hardware name: QEMU Standard PC (i440FX + PIIX, 1996), BIOS 1.15.0-1 04/01/2014
[   57.577992][ T6164] Call Trace:
[   57.578579][ T6164]  dump_stack+0x188/0x20d
[   57.579363][ T6164]  ? cgroup_apply_control_disable+0x360/0x4d0
[   57.580404][ T6164]  panic+0x2e3/0x75c
[   57.581071][ T6164]  ? add_taint.cold+0x16/0x16
[   57.581902][ T6164]  ? printk+0xba/0xed
[   57.582595][ T6164]  ? kmsg_dump_rewind_nolock+0xd9/0xd9
[   57.583656][ T6164]  ? __warn.cold+0x14/0x35
[   57.584526][ T6164]  ? __warn+0xd5/0x1c8
[   57.585272][ T6164]  ? cgroup_apply_control_disable+0x400/0x4d0
[   57.586314][ T6164]  __warn.cold+0x2f/0x35
[   57.587093][ T6164]  ? cgroup_apply_control_disable+0x400/0x4d0
[   57.588148][ T6164]  report_bug+0x28a/0x2f0
[   57.588942][ T6164]  do_error_trap+0x12b/0x220
[   57.589757][ T6164]  ? cgroup_apply_control_disable+0x400/0x4d0
[   57.590865][ T6164]  do_invalid_op+0x32/0x40
[   57.591658][ T6164]  ? cgroup_apply_control_disable+0x400/0x4d0
[   57.592855][ T6164]  invalid_op+0x23/0x30
[   57.593688][ T6164] RIP: 0010:cgroup_apply_control_disable+0x400/0x4d0
[   57.594837][ T6164] Code: 49 8d 7d 08 48 89 f8 48 c1 e8 03 0f b6 04 18 84 c0 74 08 3c 03 0f 8e b4 00 00 00 45 0f b7 6d 08 e9 6b fe ff ff e8 a0 b9 06 00 <0f> 0b e9 af fd ff ff 48 83 c4 30 5b 5d 41 5c 41 5d 41 5e 41 5f e9
[   57.598560][ T6164] RSP: 0018:ffffc90000fb7a98 EFLAGS: 00010293
[   57.599626][ T6164] RAX: ffff8880580ba080 RBX: dffffc0000000000 RCX: ffffffff816d2d4d
[   57.600997][ T6164] RDX: 0000000000000000 RSI: ffffffff816d2fa0 RDI: 0000000000000007
[   57.602423][ T6164] RBP: ffff8880699aa800 R08: ffff8880580ba080 R09: ffffed100d847104
[   57.603788][ T6164] R10: ffff88806c23881b R11: ffffed100d847103 R12: 0000000000000008
[   57.605417][ T6164] R13: 0000000000000002 R14: ffffffff89a30500 R15: ffff8880615d2000
[   57.606808][ T6164]  ? cgroup_apply_control_disable+0x1ad/0x4d0
[   57.607836][ T6164]  ? cgroup_apply_control_disable+0x400/0x4d0
[   57.608876][ T6164]  ? cgroup_apply_control_disable+0x400/0x4d0
[   57.610007][ T6164]  rebind_subsystems+0x3cd/0xb00
[   57.610878][ T6164]  ? cgroup_apply_control_disable+0x4d0/0x4d0
[   57.611964][ T6164]  ? css_populate_dir+0x288/0x450
[   57.612852][ T6164]  cgroup_setup_root+0x36a/0xac0
[   57.613767][ T6164]  ? rebind_subsystems+0xb00/0xb00
[   57.614687][ T6164]  ? init_cgroup_housekeeping+0x411/0x5a0
[   57.615771][ T6164]  cgroup1_get_tree+0xd61/0x13e3
[   57.616669][ T6164]  ? cgroup1_reconfigure+0x7d0/0x7d0
[   57.617651][ T6164]  ? security_capable+0x8f/0xc0
[   57.618565][ T6164]  vfs_get_tree+0x89/0x2f0
[   57.619342][ T6164]  do_mount+0x1315/0x1b50
[   57.620095][ T6164]  ? copy_mount_string+0x40/0x40
[   57.621001][ T6164]  ? _copy_from_user+0x13c/0x1a0
[   57.621853][ T6164]  __x64_sys_mount+0x18f/0x230
[   57.622687][ T6164]  do_syscall_64+0xf6/0x7d0
[   57.623505][ T6164]  entry_SYSCALL_64_after_hwframe+0x49/0xb3
[   57.624586][ T6164] RIP: 0033:0x4678ae
[   57.625268][ T6164] Code: 48 c7 c0 ff ff ff ff eb aa e8 0e 0e 00 00 66 2e 0f 1f 84 00 00 00 00 00 0f 1f 40 00 f3 0f 1e fa 49 89 ca b8 a5 00 00 00 0f 05 <48> 3d 01 f0 ff ff 73 01 c3 48 c7 c1 b4 ff ff ff f7 d8 64 89 01 48
[   57.628694][ T6164] RSP: 002b:00007fff6aa59de8 EFLAGS: 00000246 ORIG_RAX: 00000000000000a5
[   57.630182][ T6164] RAX: ffffffffffffffda RBX: 000000000057c988 RCX: 00000000004678ae
[   57.631562][ T6164] RDX: 00000000004d0553 RSI: 00000000004c6145 RDI: 00000000004c6108
[   57.632968][ T6164] RBP: 00000000004c6145 R08: 00000000004d5350 R09: 00007fff6aa59850
[   57.634382][ T6164] R10: 0000000000000000 R11: 0000000000000246 R12: 00000000004c6108
[   57.635822][ T6164] R13: 00000000004d0553 R14: 000000000050e8d0 R15: 0000000000000001
[   57.637521][ T6164] Dumping ftrace buffer:
[   57.638439][ T6164]    (ftrace buffer empty)
[   57.639210][ T6164] Kernel Offset: disabled
[   57.639952][ T6164] Rebooting in 1 seconds..
```

syz-crash-reporter will generate the following info:
```
Type:  , Frame: cgroup_apply_control_disable, Report: 
------------[ cut here ]------------
WARNING: CPU: 0 PID: 6164 at kernel/cgroup/cgroup.c:3111 cgroup_apply_control_disable+0x400/0x4d0 kernel/cgroup/cgroup.c:3111
Kernel panic - not syncing: panic_on_warn set ...
CPU: 0 PID: 6164 Comm: syz-executor.1 Not tainted 5.7.0 #1
Hardware name: QEMU Standard PC (i440FX + PIIX, 1996), BIOS 1.15.0-1 04/01/2014
Call Trace:
 __dump_stack lib/dump_stack.c:77 [inline]
 dump_stack+0x188/0x20d lib/dump_stack.c:118
 panic+0x2e3/0x75c kernel/panic.c:221
 __warn.cold+0x2f/0x35 kernel/panic.c:582
 report_bug+0x28a/0x2f0 lib/bug.c:195
 fixup_bug arch/x86/kernel/traps.c:175 [inline]
 fixup_bug arch/x86/kernel/traps.c:170 [inline]
 do_error_trap+0x12b/0x220 arch/x86/kernel/traps.c:267
 do_invalid_op+0x32/0x40 arch/x86/kernel/traps.c:286
 invalid_op+0x23/0x30 arch/x86/entry/entry_64.S:1027
RIP: 0010:cgroup_apply_control_disable+0x400/0x4d0 kernel/cgroup/cgroup.c:3111
Code: 49 8d 7d 08 48 89 f8 48 c1 e8 03 0f b6 04 18 84 c0 74 08 3c 03 0f 8e b4 00 00 00 45 0f b7 6d 08 e9 6b fe ff ff e8 a0 b9 06 00 <0f> 0b e9 af fd ff ff 48 83 c4 30 5b 5d 41 5c 41 5d 41 5e 41 5f e9
RSP: 0018:ffffc90000fb7a98 EFLAGS: 00010293
RAX: ffff8880580ba080 RBX: dffffc0000000000 RCX: ffffffff816d2d4d
RDX: 0000000000000000 RSI: ffffffff816d2fa0 RDI: 0000000000000007
RBP: ffff8880699aa800 R08: ffff8880580ba080 R09: ffffed100d847104
R10: ffff88806c23881b R11: ffffed100d847103 R12: 0000000000000008
R13: 0000000000000002 R14: ffffffff89a30500 R15: ffff8880615d2000
 cgroup_finalize_control kernel/cgroup/cgroup.c:3178 [inline]
 rebind_subsystems+0x3cd/0xb00 kernel/cgroup/cgroup.c:1750
 cgroup_setup_root+0x36a/0xac0 kernel/cgroup/cgroup.c:1984
 cgroup1_root_to_use kernel/cgroup/cgroup-v1.c:1190 [inline]
 cgroup1_get_tree+0xd61/0x13e3 kernel/cgroup/cgroup-v1.c:1207
 vfs_get_tree+0x89/0x2f0 fs/super.c:1547
 do_new_mount fs/namespace.c:2816 [inline]
 do_mount+0x1315/0x1b50 fs/namespace.c:3141
 __do_sys_mount fs/namespace.c:3350 [inline]
 __se_sys_mount fs/namespace.c:3327 [inline]
 __x64_sys_mount+0x18f/0x230 fs/namespace.c:3327
 do_syscall_64+0xf6/0x7d0 arch/x86/entry/common.c:295
 entry_SYSCALL_64_after_hwframe+0x49/0xb3
RIP: 0033:0x4678ae
Code: 48 c7 c0 ff ff ff ff eb aa e8 0e 0e 00 00 66 2e 0f 1f 84 00 00 00 00 00 0f 1f 40 00 f3 0f 1e fa 49 89 ca b8 a5 00 00 00 0f 05 <48> 3d 01 f0 ff ff 73 01 c3 48 c7 c1 b4 ff ff ff f7 d8 64 89 01 48
RSP: 002b:00007fff6aa59de8 EFLAGS: 00000246 ORIG_RAX: 00000000000000a5
RAX: ffffffffffffffda RBX: 000000000057c988 RCX: 00000000004678ae
RDX: 00000000004d0553 RSI: 00000000004c6145 RDI: 00000000004c6108
RBP: 00000000004c6145 R08: 00000000004d5350 R09: 00007fff6aa59850
R10: 0000000000000000 R11: 0000000000000246 R12: 00000000004c6108
R13: 00000000004d0553 R14: 000000000050e8d0 R15: 0000000000000001
Dumping ftrace buffer:
   (ftrace buffer empty)
Kernel Offset: disabled
Rebooting in 1 seconds..
```

## Usage
1. git clone syzkaller
2. cd syzkaller ; git clone syz-crash-reporter ; cd syz-crash-reporter
3. make ; make start
